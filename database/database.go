package database

import (
	"context"
	"fmt"
	"sync"

	"yiwei/database/label/filter"
	"yiwei/database/persistence"
	"yiwei/database/series"
	pb "yiwei/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Database struct {
	log *zap.Logger
	sm  map[string]*series.Series

	pb.UnimplementedDatabaseServer
	sync.RWMutex
}

func Create(log *zap.Logger) (*Database, error) {
	if err := persistence.PrepareDataDirectories(); err != nil {
		return nil, err
	}

	sl, err := persistence.ScanSeriesDirectory()
	if err != nil {
		return nil, err
	}

	sm := make(map[string]*series.Series)
	for _, sn := range sl {
		ds, err := series.Extract(sn)
		if err != nil {
			return nil, err
		}
		sm[ds.Name()] = ds
	}

	return &Database{log: log, sm: sm}, nil
}

func (db *Database) getSeries(sn string) (*series.Series, error) {
	db.Lock()
	defer db.Unlock()

	ds, ok := db.sm[sn]
	if ok {
		return ds, nil
	}

	ds, err := series.Create(sn)
	if err != nil {
		db.log.Error("failed to create series", zap.String("series", sn), zap.Error(err))
		return nil, fmt.Errorf("failed to create series")
	}

	db.log.Info("created new series", zap.String("series", ds.Name()))
	db.sm[ds.Name()] = ds
	return ds, nil
}

func (db *Database) query(req *pb.QueryRequest) (chan *pb.Entry, chan error, error) {
	db.RLock()
	defer db.RUnlock()

	ds, ok := db.sm[req.Series]
	if !ok {
		return nil, nil, fmt.Errorf("series does not exist: %s", req.Series)
	}

	c, ec := ds.Read(req.Start, req.End)

	if fe := req.Filter; fe != nil {
		f, err := filter.Parse(req.Filter)
		if err != nil {
			return nil, nil, err
		}
		c = filter.Filter(f, c)
	}

	return c, ec, nil
}

func (db *Database) Describe(_ context.Context, req *pb.DescribeRequest) (*pb.DescribeResponse, error) {
	res := &pb.DescribeResponse{}
	for sn, ds := range db.sm {
		res.Names = append(res.Names, sn)
		res.Descriptors = append(res.Descriptors, ds.Describe())
	}
	return res, nil
}

func (db *Database) Append(_ context.Context, req *pb.AppendRequest) (*pb.AppendResponse, error) {
	db.Lock()
	defer db.Unlock()

	ds, err := db.getSeries(req.Series)
	if err != nil {
		db.log.Error("failed to get series", zap.String("series", req.Series), zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to get series")
	}

	if err := ds.Append(req.Value, req.Labels); err != nil {
		db.log.Error("failed to append to series", zap.String("series", ds.Name()), zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to append into series")
	}

	return &pb.AppendResponse{}, nil
}

func (db *Database) QueryBatch(_ context.Context, req *pb.QueryRequest) (*pb.QueryBatchResponse, error) {
	c, ec, err := db.query(req)
	if err != nil {
		db.log.Error("failed to initiate query", zap.String("series", req.Series), zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to initiate query")
	}

	rl := []*pb.Reading{}
	for {
		select {
		case e, ok := <-c:
			if !ok {
				return &pb.QueryBatchResponse{Readings: rl}, nil
			}

			rl = append(rl, &pb.Reading{Entry: e})
		case err := <-ec:
			db.log.Error("query interrupted by internal error", zap.String("series", req.Series), zap.Error(err))
			return nil, status.Errorf(codes.Internal, "batch query was interrupted")
		}
	}
}

func (db *Database) QueryStream(req *pb.QueryRequest, s pb.Database_QueryStreamServer) error {
	c, ec, err := db.query(req)
	if err != nil {
		db.log.Error("failed to initiate query", zap.String("series", req.Series), zap.Error(err))
		return status.Errorf(codes.Internal, "failed to initiate query")
	}

	for {
		select {
		case e, ok := <-c:
			if !ok {
				return nil
			} else if err := s.Send(&pb.Reading{Entry: e}); err != nil {
				db.log.Error("query interrupted while trying to deliver reading", zap.String("series", req.Series), zap.Error(err))
				return status.Errorf(codes.Internal, "failed to deliver query stream")
			}
		case err := <-ec:
			db.log.Error("query interrupted by internal error", zap.String("series", req.Series), zap.Error(err))
			return status.Errorf(codes.Internal, "stream query was interrupted")
		}
	}
}

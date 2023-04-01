package database

import (
	"context"
	"fmt"
	"sync"

	"yiwei/database/persistence"
	"yiwei/database/series"
	pb "yiwei/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Database struct {
	log *zap.Logger
	rw  sync.RWMutex
	sm  map[string]*series.Series

	pb.UnimplementedDatabaseServer
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

func (db *Database) CreateSeries(ctxt context.Context, req *pb.CreateSeriesRequest) (*pb.CreateSeriesResponse, error) {
	db.rw.Lock()
	defer db.rw.Unlock()

	if _, ok := db.sm[req.Name]; ok {
		return &pb.CreateSeriesResponse{}, nil
	}

	ds, err := series.Create(req.Name, req.IndexGenerator, req.ContextLabels)
	if err != nil {
		db.log.Error("failed to create series", zap.String("series", ds.Name()), zap.Error(err))
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to create series"))
	}

	db.log.Info("created new series", zap.String("series", ds.Name()))
	db.sm[ds.Name()] = ds
	return &pb.CreateSeriesResponse{}, nil
}

func (db *Database) Append(ctxt context.Context, req *pb.AppendRequest) (*pb.AppendResponse, error) {
	ds, ok := db.sm[req.Series]
	if !ok {
		db.log.Error("intended to append to non-existent series", zap.String("series", req.Series))
		return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("series does not exist"))
	}

	if err := ds.Append(req.Value, req.Labels); err != nil {
		db.log.Error("failed to append to series", zap.String("series", ds.Name()), zap.Error(err))
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to append into series"))
	}

	return &pb.AppendResponse{}, nil
}

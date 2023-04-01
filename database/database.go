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
	sm  map[string]*series.Series
	rw  sync.RWMutex

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

func (db *Database) getSeries(sn string) (*series.Series, error) {
	db.rw.Lock()
	defer db.rw.Unlock()

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

func (db *Database) Append(ctxt context.Context, req *pb.AppendRequest) (*pb.AppendResponse, error) {
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

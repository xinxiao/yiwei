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

	rw sync.RWMutex

	pb.UnimplementedDatabaseServer
}

func Create(log *zap.Logger) (*Database, error) {

	if err := persistence.PrepareDataDirectories(); err != nil {
		return nil, err
	}

	return &Database{log: log, sm: make(map[string]*series.Series)}, nil
}

func (db *Database) CreateSeries(ctxt context.Context, req *pb.CreateSeriesRequest) (*pb.CreateSeriesResponse, error) {
	db.rw.Lock()
	defer db.rw.Unlock()

	if _, ok := db.sm[req.Name]; ok {
		return &pb.CreateSeriesResponse{}, nil
	}

	ds, err := series.Create(req.Name, req.IndexGenerator, req.ContextLabels)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to create series %s: %s", req.Name, err))
	}

	db.sm[ds.Name()] = ds
	return &pb.CreateSeriesResponse{}, nil
}

func (db *Database) Append(ctxt context.Context, req *pb.AppendRequest) (*pb.AppendResponse, error) {
	ds, ok := db.sm[req.Series]
	if !ok {
		return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("series does not exist: %s", req.Series))
	}

	if err := ds.Append(req.Value, req.Labels); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to append into series %s: %s", ds.Name(), err))
	}

	return &pb.AppendResponse{}, nil
}

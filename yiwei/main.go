package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"yiwei/database"
	pb "yiwei/proto"

	"go.uber.org/zap"
)

var (
	prod = flag.Bool("prod", false, "is prod mode")
)

func main() {
	flag.Parse()

	log := CreateLogger(*prod)

	db, err := database.Create(log)
	if err != nil {
		log.Fatal("failed to create database", zap.Error(err))
	}

	svr := CreateServer(log)
	pb.RegisterDatabaseServer(svr, db)

	ec := make(chan error, 1)
	sc := make(chan os.Signal, 1)

	go Launch(log, svr, ec)
	signal.Notify(sc, syscall.SIGTERM, syscall.SIGINT)

	select {
	case err := <-ec:
		log.Fatal("failed to launch grpc server", zap.Error(err))
	case <-sc:
		log.Info("stopping grpc server")
	}
}

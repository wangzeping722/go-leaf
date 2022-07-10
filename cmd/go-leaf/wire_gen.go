// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go-leaf/internal/biz"
	"go-leaf/internal/conf"
	"go-leaf/internal/data"
	"go-leaf/internal/server"
	"go-leaf/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, leaf *conf.Leaf, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(leaf, logger)
	if err != nil {
		return nil, nil, err
	}
	leafAllocRepo := data.NewLeafAllocRepo(dataData, logger)
	leafAllocUsecase := biz.NewLeafAllocUsecase(leafAllocRepo, logger)
	segmentService := service.NewSegmentService(leaf, leafAllocUsecase, logger)
	leafService := service.NewLeafService(segmentService, logger)
	httpServer := server.NewHTTPServer(confServer, leafService, logger)
	grpcServer := server.NewGRPCServer(confServer, leafService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}

// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/adshao/ordinals-indexer/internal/biz"
	"github.com/adshao/ordinals-indexer/internal/conf"
	"github.com/adshao/ordinals-indexer/internal/data"
	"github.com/adshao/ordinals-indexer/internal/server"
	"github.com/adshao/ordinals-indexer/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	collectionRepo := data.NewCollectionRepo(dataData, logger)
	collectionUsecase := biz.NewCollectionUsecase(collectionRepo, logger)
	collectionService := service.NewCollectionService(collectionUsecase, logger)
	tokenRepo := data.NewTokenRepo(dataData, logger)
	tokenUsecase := biz.NewTokenUsecase(tokenRepo, logger)
	tokenService := service.NewTokenService(tokenUsecase, logger)
	inscriptionRepo := data.NewInscriptionRepo(dataData, logger)
	inscriptionUsecase := biz.NewInscriptionUsecase(inscriptionRepo, logger)
	inscriptionService := service.NewInscriptionService(inscriptionUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, collectionService, tokenService, inscriptionService, logger)
	httpServer := server.NewHTTPServer(confServer, collectionService, tokenService, inscriptionService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}

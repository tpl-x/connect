// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tpl-x/conec/internal/config"
	"github.com/tpl-x/conec/internal/server"
	"github.com/tpl-x/conec/internal/service"
	"github.com/tpl-x/conec/pkg/logger"
)

// Injectors from wire.go:

// wireApp init for builder backend
func wireApp() (*app, error) {
	serverConfig, err := config.LoadDefaultConfig()
	if err != nil {
		return nil, err
	}
	logConfig := serverConfig.Log
	zapLogger := logger.NewZapLogger(logConfig)
	pingService := service.NewPingService(zapLogger)
	pingServer := server.NewPingServer(serverConfig, zapLogger, pingService)
	mainApp := newApp(pingServer)
	return mainApp, nil
}

// wire.go:

var appSet = wire.NewSet(config.LoadDefaultConfig, wire.FieldsOf(new(*config.ServerConfig), "Log"), logger.NewZapLogger)
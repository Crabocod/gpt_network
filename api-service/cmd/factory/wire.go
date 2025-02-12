//go:build wireinject
// +build wireinject

package factory

import (
	"github.com/Crabocod/gpt_network/api-service/internal/app/apiserver"
	"github.com/Crabocod/gpt_network/api-service/internal/app/grpcserver"
	"github.com/Crabocod/gpt_network/api-service/internal/config"
	"github.com/Crabocod/gpt_network/api-service/internal/db"
	"github.com/google/wire"
)

func InitializeApp() (*App, error) {
	wire.Build(
		config.Load,
		db.Connect,
		apiserver.New,
		grpcserver.New,
		New,
	)
	return &App{}, nil
}

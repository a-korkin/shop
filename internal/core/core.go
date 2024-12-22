package core

import (
	"github.com/a-korkin/shop/internal/ports"
)

type AppState struct {
	DbConn   ports.DbConnect
	ApiPort  string
	GrpcPort string
}

func NewAppState(dbConn ports.DbConnect, apiPort string, grpcPort string) AppState {
	return AppState{
		DbConn:   dbConn,
		ApiPort:  apiPort,
		GrpcPort: grpcPort,
	}
}

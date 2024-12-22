package core

import (
	"github.com/a-korkin/shop/internal/ports"
)

type AppState struct {
	DbConn   ports.DbConnect
	GrpcPort string
}

func NewAppState(dbConn ports.DbConnect, port string) *AppState {
	return &AppState{
		DbConn:   dbConn,
		GrpcPort: port,
	}
}

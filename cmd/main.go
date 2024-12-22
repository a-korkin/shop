package main

import (
	"fmt"
	"github.com/a-korkin/shop/configs"
	"github.com/a-korkin/shop/internal/api"
	// "github.com/a-korkin/shop/internal/rpc"
	"github.com/a-korkin/shop/internal/adapters"
	"log"
)

func main() {
	dbConn, err := adapters.NewDBConnect(configs.GetDBConnection())
	if err != nil {
		log.Fatalf("failed to create connection to db: %s", err)
	}
	defer func() {
		if err := dbConn.Db.Close(); err != nil {
			log.Fatalf("failed to close connection to db: %s", err)
		}
	}()
	apiPort := configs.GetWebApiPort()
	api.Run(fmt.Sprintf(":%s", apiPort), dbConn)

	// grpcPort := configs.GetGrpcPort()
	// server := rpc.NewShopServer()
	// server.Run(fmt.Sprintf(":%s", grpcPort))
}

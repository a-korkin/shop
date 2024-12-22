package main

import (
	"fmt"
	"log"
	"os"

	"github.com/a-korkin/shop/configs"
	"github.com/a-korkin/shop/internal/adapters"
	"github.com/a-korkin/shop/internal/api"
	"github.com/a-korkin/shop/internal/core"
	"github.com/a-korkin/shop/internal/rpc"
)

func help() {
	fmt.Printf(`Params: 
	-a, --api	run web api
	-g, --grpc	run grpc server
	-h, --help	show help
`)
}

func runWebApi(appState *core.AppState) {
	log.Printf("web api running on port: %s", appState.ApiPort)
	api.Run(appState)
}

func runGrpcServer(appState *core.AppState) {
	server := rpc.NewShopServer(appState)
	log.Printf("grpc server running on port: %s", appState.GrpcPort)
	server.Run(appState.GrpcPort)
}

func main() {
	if len(os.Args) == 1 || len(os.Args) > 2 {
		fmt.Printf("You must choice which app to run. See help.\n")
		help()
		os.Exit(1)
	}
	dbConn, err := adapters.NewDBConnect(configs.GetDBConnection())
	if err != nil {
		log.Fatalf("failed to create connection to db: %s", err)
	}
	defer func() {
		if err := dbConn.Db.Close(); err != nil {
			log.Fatalf("failed to close connection to db: %s", err)
		}
	}()
	appState := core.NewAppState(
		dbConn,
		fmt.Sprintf(":%s", configs.GetWebApiPort()),
		fmt.Sprintf(":%s", configs.GetGrpcPort()))
	switch os.Args[1] {
	case "-a":
		runWebApi(&appState)
	case "--api":
		runWebApi(&appState)
	case "-g":
		runGrpcServer(&appState)
	case "--grpc":
		runGrpcServer(&appState)
	case "-h":
		help()
	case "--help":
		help()
	default:
		fmt.Printf("Wrong argument: %s\n", os.Args[1])
		help()
		os.Exit(1)
	}
}

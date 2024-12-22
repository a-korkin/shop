package main

import (
	// "fmt"
	// "github.com/a-korkin/shop/configs"
	// "github.com/a-korkin/shop/internal/api"
	"github.com/a-korkin/shop/internal/rpc"
)

func main() {
	// port := configs.GetWebApiPort()
	// api.Run(fmt.Sprintf(":%s", port))

	server := rpc.NewShopServer()
	server.Run()
}

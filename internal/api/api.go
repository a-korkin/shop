package api

import (
	"context"
	"github.com/a-korkin/shop/internal/api/handlers"
	pb "github.com/a-korkin/shop/internal/common"
	"github.com/a-korkin/shop/internal/core"
	"github.com/a-korkin/shop/internal/ports"
	"github.com/a-korkin/shop/internal/rpc"
	"github.com/a-korkin/shop/internal/tools"
	"net/http"
)

type ShopHandler struct {
	Db         ports.DbConnect
	GrpcClient pb.ShopServiceClient
}

func (h *ShopHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resource := tools.GetResource(r.RequestURI)
	ctx := context.Background()
	switch resource {
	case "items":
		handlers.ItemHandle(ctx, r.RequestURI, h.GrpcClient, w, r)
	case "users":
		handlers.UserHandle(ctx, r.RequestURI, h.GrpcClient, w, r)
	default:
		http.Error(w, "resource not found", http.StatusNotFound)
	}
}

func Run(appState *core.AppState) error {
	server := http.Server{
		Addr: appState.ApiPort,
	}
	grpcClient, err := rpc.NewGrpcClient(appState.GrpcPort)
	if err != nil {
		return err
	}
	handler := ShopHandler{
		Db:         appState.DbConn,
		GrpcClient: grpcClient,
	}
	http.Handle("/", &handler)
	if err := server.ListenAndServe(); err != nil {
		return err
	}
	defer server.Close()
	return nil
}

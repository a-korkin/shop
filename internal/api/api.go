package api

import (
	"encoding/json"
	"log"
	"net/http"

	"context"
	pb "github.com/a-korkin/shop/internal/common"
	"github.com/a-korkin/shop/internal/core"
	"github.com/a-korkin/shop/internal/ports"
	"github.com/a-korkin/shop/internal/rpc"
	"github.com/a-korkin/shop/internal/tools"
)

type ShopHandler struct {
	Db         ports.DbConnect
	GrpcPort   string
	GrpcClient pb.ShopServiceClient
}

func (h *ShopHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resource := tools.GetResource(r.RequestURI)

	switch resource {
	case "items":
		id, err := tools.GetId(r.RequestURI)
		if err == nil {
			item, err := h.GrpcClient.GetItem(
				context.Background(), &pb.ItemId{Id: int32(id)})
			if err != nil {
				log.Fatalf("failed to get item: %s", err)
			}
			if item.Id == 0 {
				http.Error(w, "item not found", http.StatusNotFound)
				return
			}
			encoder := json.NewEncoder(w)
			err = encoder.Encode(item)
			if err != nil {
				log.Fatalf("failed to encode item: %s", err)
			}
		}
	default:
		http.Error(w, "resource not found", http.StatusNotFound)
	}
}

func Run(appState *core.AppState) {
	server := http.Server{
		Addr: appState.ApiPort,
	}
	grpcClient, err := rpc.NewGrpcClient(appState.GrpcPort)
	if err != nil {
		log.Fatalf("failed to create client: %s", err)
	}
	handler := ShopHandler{
		Db:         appState.DbConn,
		GrpcPort:   appState.GrpcPort,
		GrpcClient: grpcClient,
	}
	http.Handle("/", &handler)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to start web api: %s", err)
	}
	defer server.Close()
}

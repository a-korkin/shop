package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"context"
	pb "github.com/a-korkin/shop/internal/common"
	"github.com/a-korkin/shop/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ShopHandler struct {
	Db ports.DbConnect
}

func callGrpc() (*pb.Item, error) {
	conn, err := grpc.NewClient(
		"localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create grpc client: %s", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalf("failed to close grpc connection: %s", err)
		}
	}()
	client := pb.NewShopServiceClient(conn)
	item, err := client.GetItem(context.Background(), &pb.ItemId{Id: 1})
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (h *ShopHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resource := strings.Split(strings.ToLower(r.RequestURI), "/")[1]

	switch resource {
	case "items":
		// item, err := h.Db.GetItem(1)
		item, err := callGrpc()
		if err != nil {
			log.Fatalf("failed to get item: %s", err)
		}
		encoder := json.NewEncoder(w)
		err = encoder.Encode(item)
		if err != nil {
			log.Fatalf("failed to encode item: %s", err)
		}
	}
}

func Run(port string, dbConn ports.DbConnect) {
	server := http.Server{
		Addr: port,
	}
	handler := ShopHandler{
		Db: dbConn,
	}
	http.Handle("/", &handler)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to start web api: %s", err)
	}
	defer server.Close()
}

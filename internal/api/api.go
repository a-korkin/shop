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
	GrpcClient pb.ShopServiceClient
}

func (h *ShopHandler) itemsHandler(uri string, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id, err := tools.GetId(uri)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if id >= 0 {
			item, err := h.GrpcClient.GetItem(
				context.Background(), &pb.ItemId{Id: int32(id)})
			if err != nil {
				log.Fatalf("failed to get item: %s", err)
			}
			if item.Id <= 0 {
				http.Error(w, "item not found", http.StatusNotFound)
				return
			}
			encoder := json.NewEncoder(w)
			err = encoder.Encode(item)
			if err != nil {
				log.Fatalf("failed to encode item: %s", err)
			}
		} else {
			pageParams := tools.GetPageParams(r.URL.RawQuery)
			items, err := h.GrpcClient.GetItems(context.Background(), pageParams)
			if err != nil {
				log.Fatalf("failed to get items: %s", err)
			}
			encoder := json.NewEncoder(w)
			if err := encoder.Encode(&items); err != nil {
				log.Fatalf("failed to encode list items: %s", err)
			}
		}
	case "POST":
		decoder := json.NewDecoder(r.Body)
		in := pb.ItemDto{}
		if err := decoder.Decode(&in); err != nil {
			log.Fatalf("failed to unmarshalling item: %s", err)
		}
		item, err := h.GrpcClient.CreateItem(context.Background(), &in)
		if err != nil {
			log.Fatalf("failed to create item: %s", err)
		}
		w.WriteHeader(http.StatusCreated)
		encoder := json.NewEncoder(w)
		encoder.Encode(&item)
	case "PUT":
		id, err := tools.GetId(uri)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if id <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		decoder := json.NewDecoder(r.Body)
		in := pb.ItemDto{}
		if err := decoder.Decode(&in); err != nil {
			log.Fatalf("failed to update item: %s", err)
		}
		item := pb.Item{
			Id:    int32(id),
			Title: in.Title,
			Price: in.Price,
		}
		_, err = h.GrpcClient.UpdItem(context.Background(), &item)
		if err != nil {
			log.Fatalf("failed to update item: %s", err)
		}
		w.WriteHeader(http.StatusOK)
		encoder := json.NewEncoder(w)
		encoder.Encode(&item)
	case "DELETE":
		id, err := tools.GetId(uri)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if id <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = h.GrpcClient.DropItem(context.Background(), &pb.ItemId{Id: int32(id)})
		if err != nil {
			log.Fatalf("failed to delete item: %s", err)
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func (h *ShopHandler) usersHandler(uri string, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		userIn := pb.UserDto{}
		err := json.NewDecoder(r.Body).Decode(&userIn)
		if err != nil {
			log.Fatalf("failed to create user: %s", err)
		}
		userOut, err := h.GrpcClient.CreateUser(context.Background(), &userIn)
		if err != nil {
			log.Fatalf("failed to creating user: %s", err)
		}
		json.NewEncoder(w).Encode(userOut)
	}
}

func (h *ShopHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resource := tools.GetResource(r.RequestURI)
	switch resource {
	case "items":
		h.itemsHandler(r.RequestURI, w, r)
	case "users":
		h.usersHandler(r.RequestURI, w, r)
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

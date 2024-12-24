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
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (h *ShopHandler) usersHandler(uri string, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id, err := tools.GetId(uri)
		if err != nil {
			log.Fatalf("failed to parse id: %s", err)
		}
		if id < 0 {
			pageParams := tools.GetPageParams(r.URL.RawQuery)
			userStream, err := h.GrpcClient.GetUsers(context.Background(), pageParams)
			if err != nil {
				log.Fatalf("failed to get users: %s", err)
			}
			users := make([]*pb.User, 0)
			for {
				user, err := userStream.Recv()
				if err != nil {
					break
				}
				users = append(users, user)
			}
			if err = json.NewEncoder(w).Encode(users); err != nil {
				log.Fatalf("failed to marshalling list of users: %s", err)
			}
		} else {
			user, err := h.GrpcClient.GetUser(
				context.Background(), &pb.UserId{Id: int32(id)})
			if err != nil {
				log.Fatalf("failed to get user: %s", err)
			}
			if err = json.NewEncoder(w).Encode(&user); err != nil {
				log.Fatalf("failed to marshalling user: %s", err)
			}
		}
	case "POST":
		userIn := pb.UserDto{}
		err := json.NewDecoder(r.Body).Decode(&userIn)
		if err != nil {
			log.Fatalf("failed to unmarshalling user: %s", err)
		}
		userOut, err := h.GrpcClient.CreateUser(context.Background(), &userIn)
		if err != nil {
			log.Fatalf("failed to creating user: %s", err)
		}
		if err := json.NewEncoder(w).Encode(userOut); err != nil {
			log.Fatalf("failed to marshalling user: %s", err)
		}
	case "PUT":
		id, err := tools.GetId(uri)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		userIn := pb.UserDto{}
		err = json.NewDecoder(r.Body).Decode(&userIn)
		if err != nil {
			log.Fatalf("failed to unmarshalling user: %s", err)
		}
		user := pb.User{
			Id:        int32(id),
			LastName:  userIn.LastName,
			FirstName: userIn.FirstName,
		}
		log.Printf("user: %v", &user)
		userOut, err := h.GrpcClient.UpdUser(context.Background(), &user)
		if err != nil {
			log.Fatalf("failed to update user: %s", err)
		}
		if err = json.NewEncoder(w).Encode(&userOut); err != nil {
			log.Fatalf("failed to marshalling user: %s", err)
		}
	case "DELETE":
		id, err := tools.GetId(uri)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = h.GrpcClient.DropUser(context.Background(), &pb.UserId{Id: int32(id)})
		if err != nil {
			log.Fatalf("failed to delete user: %s", err)
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
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

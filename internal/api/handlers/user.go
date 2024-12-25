package handlers

import (
	"context"
	"encoding/json"
	pb "github.com/a-korkin/shop/internal/common"
	"github.com/a-korkin/shop/internal/tools"
	"log"
	"net/http"
)

func UserHandle(
	ctx context.Context, uri string,
	client pb.ShopServiceClient, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id, err := tools.GetId(uri)
		if err != nil {
			log.Fatalf("failed to parse id: %s", err)
		}
		if id < 0 {
			pageParams := tools.GetPageParams(r.URL.RawQuery)
			userStream, err := client.GetUsers(ctx, pageParams)
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
			user, err := client.GetUser(ctx, &pb.UserId{Id: int32(id)})
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
		userOut, err := client.CreateUser(ctx, &userIn)
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
		userOut, err := client.UpdUser(ctx, &user)
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
		_, err = client.DropUser(ctx, &pb.UserId{Id: int32(id)})
		if err != nil {
			log.Fatalf("failed to delete user: %s", err)
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

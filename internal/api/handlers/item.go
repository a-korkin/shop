package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	pb "github.com/a-korkin/shop/internal/common"
	"github.com/a-korkin/shop/internal/tools"
)

func ItemHandle(
	ctx context.Context, client pb.ShopServiceClient,
	w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id, err := tools.GetId(r.RequestURI)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if id >= 0 {
			item, err := client.GetItem(ctx, &pb.ItemId{Id: int32(id)})
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
			items, err := client.GetItems(ctx, pageParams)
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
		item, err := client.CreateItem(ctx, &in)
		if err != nil {
			log.Fatalf("failed to create item: %s", err)
		}
		w.WriteHeader(http.StatusCreated)
		encoder := json.NewEncoder(w)
		encoder.Encode(&item)
	case "PUT":
		id, err := tools.GetId(r.RequestURI)
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
		_, err = client.UpdItem(ctx, &item)
		if err != nil {
			log.Fatalf("failed to update item: %s", err)
		}
		w.WriteHeader(http.StatusOK)
		encoder := json.NewEncoder(w)
		encoder.Encode(&item)
	case "DELETE":
		id, err := tools.GetId(r.RequestURI)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if id <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = client.DropItem(ctx, &pb.ItemId{Id: int32(id)})
		if err != nil {
			log.Fatalf("failed to delete item: %s", err)
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

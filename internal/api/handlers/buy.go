package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	pb "github.com/a-korkin/shop/internal/common"
)

type Purchase struct {
	UserId         int32
	ItemId         int32
	TimeOfPurchase time.Time
	CountItems     int32
}

func BuyHandle(
	ctx context.Context, client pb.ShopServiceClient,
	w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		in := pb.PurchaseDto{}
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			log.Fatalf("failed to unmarshalling purchase: %s", err)
		}
		out, err := client.Buy(ctx, &in)
		if err != nil {
			log.Fatalf("failed to buy items: %s", err)
		}
		p := Purchase{
			UserId:         out.UserId,
			ItemId:         out.ItemId,
			TimeOfPurchase: out.TimeOfPurchase.AsTime(),
			CountItems:     out.CountItems,
		}
		if err = json.NewEncoder(w).Encode(&p); err != nil {
			log.Fatalf("failed to marshalling purchase: %s", err)
		}
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

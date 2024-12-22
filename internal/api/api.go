package api

import (
	"encoding/json"
	"github.com/a-korkin/shop/internal/ports"
	"log"
	"net/http"
	"strings"
)

type ShopHandler struct {
	Db ports.DbConnect
}

func (h *ShopHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resource := strings.Split(strings.ToLower(r.RequestURI), "/")[1]

	switch resource {
	case "items":
		item, err := h.Db.GetItem(1)
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

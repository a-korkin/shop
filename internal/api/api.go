package api

import (
	"github.com/a-korkin/shop/internal/ports"
	"log"
	"net/http"
	"strings"
)

type ShopHandler struct{}

func (h *ShopHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resource := strings.Split(strings.ToLower(r.RequestURI), "/")[1]
	router(resource)
	w.Write([]byte("hello from api"))
}

func router(resource string) {
	switch resource {
	case "items":

	}
}

func Run(port string, dbConn ports.DbConnect) {
	server := http.Server{
		Addr: port,
	}
	handler := ShopHandler{}
	http.Handle("/", &handler)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to start web api: %s", err)
	}
	defer server.Close()
}

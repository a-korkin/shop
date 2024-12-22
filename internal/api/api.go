package api

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from api"))
}

func Run() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", handler)
	server.ListenAndServe()
}

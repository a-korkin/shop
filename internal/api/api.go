package api

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from api"))
}

func Run(port string) {
	server := http.Server{
		Addr: port,
	}
	http.HandleFunc("/", handler)
	server.ListenAndServe()
}

package server

import (
	"log"
	"net/http"
	"os"

	"github.com/RedWing96399/shortscreen/routes"
)

func Run(mux *http.ServeMux) {

	s := http.Server{
		Addr:    os.Getenv("SERVER_ADDR"),
		Handler: mux,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func GetMux() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", routes.Home)

	return mux
}

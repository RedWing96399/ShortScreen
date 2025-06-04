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
	mux.HandleFunc("GET /images", routes.ListImages)
	mux.HandleFunc("GET /images/{id}", routes.GetImage)
	mux.HandleFunc("POST /images", routes.UploadImages)
	mux.HandleFunc("PATCH /images/{id}", routes.UpdateImage)
	mux.HandleFunc("DELTE /images/{id}", routes.DeleteImage)

	return mux
}

package routes

import (
	"fmt"
	"net/http"
)

func UploadImages(w http.ResponseWriter, r *http.Request) {

	// this handler uploads one or more images (through the request payload) to the server.
	// the ownership of these images belongs to the one who uploaded these images.

	fmt.Fprintf(w, "this handler uploads images to the server")
}

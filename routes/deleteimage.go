package routes

import (
	"fmt"
	"net/http"
)

func DeleteImage(w http.ResponseWriter, r *http.Request) {

	imageId := r.PathValue("id")
	fmt.Fprintf(w, "this handler deletes the said image :%s", imageId)
}

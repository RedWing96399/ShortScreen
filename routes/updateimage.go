package routes

import (
	"fmt"
	"net/http"
)

func UpdateImage(w http.ResponseWriter, r *http.Request) {

	// this function updates the meta data of the image.
	//  the image id is from the path parameter and the new data should be fetched from request payload.
	imageId := r.PathValue("id")

	fmt.Fprintf(w, "this is update image handler %s", imageId)
}

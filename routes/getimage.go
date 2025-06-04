package routes

import (
	"fmt"
	"net/http"
)

func GetImage(w http.ResponseWriter, r *http.Request) {

	// the imageId should be used to call an underlying service layer function to get image data.
	// the data thus obtained, should be send in the response body.

	// image id should be validated before use.
	imageId := r.PathValue("id")

	// this method requires that the user is authenticated.
	fmt.Fprintf(w, "this function returns the image requested by the user, which is :%s", imageId)
}

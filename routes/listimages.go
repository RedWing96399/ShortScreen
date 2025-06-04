package routes

import (
	"fmt"
	"net/http"
)

func ListImages(w http.ResponseWriter, r *http.Request) {

	// the  authenticated user's data is the one that is fetched.
	// no data should be displayed for the user if the data requested is not their's
	// this method requires that the user is authenticated.
	fmt.Fprintf(w, "this function returns the list of images owned by the user")
}

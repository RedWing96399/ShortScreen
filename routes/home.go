package routes

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	respString := fmt.Sprintf(
		"this is the home function\nRequest Method:%s\nRequest Pattern:%s\nRequest URI:%s",
		r.Method, r.Pattern, r.RequestURI)

	w.WriteHeader(http.StatusAccepted)
	w.Header().Add("roman", "demidov")
	fmt.Fprintln(w, respString)
}

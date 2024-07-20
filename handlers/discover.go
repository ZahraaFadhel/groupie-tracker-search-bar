package groupieTracker

import (
	"fmt"
	"net/http"
)

func DiscoverHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("discover error")
	ErrorHandler(w, r, http.StatusNotFound)
}

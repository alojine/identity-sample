package routes

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/heartbeat", heartbeatHandler)

	return mux
}

func heartbeatHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "live!")
}

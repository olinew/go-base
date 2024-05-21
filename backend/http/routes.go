package http

import (
	"net/http"
)

func CoreRoutes() http.Handler {
	core := http.NewServeMux()
	core.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	return core
}

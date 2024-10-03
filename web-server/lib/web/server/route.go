package server

import (
	"net/http"
	"stock/lib/web/handler"
)

func route() *http.ServeMux {
	h := handler.New(handler.Option{})
	mux := http.NewServeMux()

	mux.HandleFunc("/", h.Home)

	return mux
}

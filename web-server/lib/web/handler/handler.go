package handler

import "net/http"

type I interface {
	Home(w http.ResponseWriter, r *http.Request)
}

type Option struct {
}

type Handler struct {
}

func New(opt Option) I {
	return &Handler{}
}

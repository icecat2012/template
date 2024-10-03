package handler

import (
	"fmt"
	"net/http"
)

func (p *Handler) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

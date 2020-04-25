package handler

import "fmt"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello golang Voronezh!")
}
package handler

import "github.com/gorilla/mux"

func (h *Handler) ApiRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/leagues", h.Leagues)
	router.HandleFunc("/teams", h.Teams)

	return router
}

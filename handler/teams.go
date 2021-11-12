package handler

import "net/http"

func (h *Handler) Teams(w http.ResponseWriter, r *http.Request) {
	teams, err := h.fetcher.Teams(r.Context())
	if err != nil {
		errorStatusHTTP(w, err)

		return
	}

	httpResponseWithBytes(w, teams)
}

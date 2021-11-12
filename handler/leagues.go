package handler

import (
	"net/http"
)

func (h *Handler) Leagues(w http.ResponseWriter, r *http.Request) {
	leagues, err := h.fetcher.Leagues(r.Context())
	if err != nil {
		errorStatusHTTP(w, err)

		return
	}

	httpResponseWithBytes(w, leagues)
}

package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/serhiihuberniuk/unstable-api/models"
)

func httpResponseWithBytes(w http.ResponseWriter, content interface{}) {
	err := json.NewEncoder(w).Encode(content)
	if err != nil {
		errorStatusHTTP(w, err)

		return
	}
}

func errorStatusHTTP(w http.ResponseWriter, err error) {
	log.Println(err)

	if errors.Is(err, models.ErrTimeout) {
		http.Error(w, http.StatusText(http.StatusRequestTimeout), http.StatusRequestTimeout)

		return
	}

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

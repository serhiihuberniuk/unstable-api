package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/serhiihuberniuk/unstable-api/models"
)

func httpResponseWithBytes(w http.ResponseWriter, content interface{}) {
	data, err := json.Marshal(content)
	if err != nil {
		errorStatusHTTP(w, err)

		return
	}

	_, err = w.Write(data)
	if err != nil {
		errorStatusHTTP(w, err)
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

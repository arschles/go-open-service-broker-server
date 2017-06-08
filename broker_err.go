package osbserver

import (
	"encoding/json"
	"net/http"
)

type brokerErr struct {
	Description string `json:"description"`
}

func internalServerError(w http.ResponseWriter, err error) {
	json.NewEncoder(w).Encode(brokerErr{Description: err.Error()})
}

package osbserver

import (
	"encoding/json"
	"net/http"

	"github.com/arschles/osbserver/request"
	"github.com/gorilla/mux"
)

func bind(ops Operations) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		instID, ok := mux.Vars(r)["instance_id"]
		if !ok {
			http.Error(w, "missing instance ID", http.StatusBadRequest)
			return
		}
		bindID, ok := mux.Vars(r)["binding_id"]
		if !ok {
			http.Error(w, "missing binding ID", http.StatusBadRequest)
			return
		}
		bindReqBody := new(request.Bind)
		if err := json.NewDecoder(r.Body).Decode(bindReqBody); err != nil {
			logger.Printf("error decoding body (%s)", err)
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		bindRes, err := ops.Bind(instID, bindID, bindReqBody)
		if err != nil {
			writeBrokerErr(w, err)
			return
		}
		if err := json.NewEncoder(w).Encode(bindRes); err != nil {
			logger.Printf("error encoding JSON (%s)", err)
			http.Error(w, "error encoding JSON", http.StatusInternalServerError)
			return
		}
	}
}

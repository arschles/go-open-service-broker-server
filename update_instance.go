package osbserver

import (
	"encoding/json"
	"net/http"

	"github.com/arschles/osbserver/request"
	"github.com/gorilla/mux"
)

func updateInstance(ops Operations) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		instID, ok := mux.Vars(r)["instance_id"]
		if !ok {
			http.Error(w, "missing instance ID", http.StatusBadRequest)
			return
		}
		reqBody := new(request.UpdateServiceInstance)
		if err := json.NewDecoder(r.Body).Decode(reqBody); err != nil {
			logger.Printf("bad JSON request body (%s)", err)
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		resp, err := ops.UpdateInstance(instID, reqBody)
		if err != nil {
			writeBrokerErr(w, err)
			return
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			logger.Printf("error encoding JSON (%s)", err)
			http.Error(w, "error encoding JSON", http.StatusInternalServerError)
			return
		}
	}
}

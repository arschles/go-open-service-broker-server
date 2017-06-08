package osbserver

import (
	"encoding/json"
	"net/http"

	"github.com/arschles/osbserver/request"
	"github.com/gorilla/mux"
)

func provision(ops Operations) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		instID, ok := mux.Vars(r)["instance_id"]
		if !ok {
			http.Error(w, "missing instance ID", http.StatusBadRequest)
			return
		}
		provReq := new(request.Provision)
		if err := json.NewDecoder(r.Body).Decode(provReq); err != nil {
			logger.Printf("invalid request body (%s)", err)
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		provResp, err := ops.Provision(instID, provReq)
		if err != nil {
			writeBrokerErr(w, err)
			return
		}
		if err := json.NewEncoder(w).Encode(provResp); err != nil {
			logger.Printf("error encoding json (%s)", err)
			http.Error(w, "error encoding json", http.StatusInternalServerError)
			return
		}
	}
}

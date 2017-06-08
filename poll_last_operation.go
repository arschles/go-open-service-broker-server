package osbserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func pollLastOperation(ops Operations) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		instID, ok := mux.Vars(r)["instance_id"]
		if !ok {
			http.Error(w, "missing instance ID", http.StatusBadRequest)
			return
		}

		svcID := r.URL.Query().Get("service_id")
		planID := r.URL.Query().Get("plan_id")
		op := r.URL.Query().Get("operation")
		resp, err := ops.PollLastOperation(instID, &PollLastOperationParams{
			ServiceID: svcID,
			PlanID:    planID,
			Operation: op,
		})
		if err != nil {
			writeBrokerErr(w, err)
			return
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, fmt.Sprintf("error encoding json (%s)", err), http.StatusInternalServerError)
			return
		}
	}
}

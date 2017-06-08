package osbserver

import (
	"encoding/json"
	"net/http"
)

func getCatalog(ops Operations) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cat, err := ops.Catalog()
		if err != nil {
			writeBrokerErr(w, err)
			return
		}
		if err := json.NewEncoder(w).Encode(cat); err != nil {
			logger.Printf("error encoding JSON (%s)", err)
			http.Error(w, "error encoding JSON", http.StatusInternalServerError)
			return
		}
	}
}

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
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

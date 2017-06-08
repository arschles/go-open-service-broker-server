package osbserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Run(port int, operations Operations) error {
	r := mux.NewRouter()

	// catalog
	r.HandleFunc("/v2/catalog", getCatalog(operations)).Methods("GET")

	// poll last operation
	r.HandleFunc("/v2/service_instances/{instance_id}/last_operation", pollLastOperation(operations)).Methods("GET")

	hostStr := fmt.Sprintf(":%d", port)
	http.ListenAndServe(hostStr, r)
	return nil
}

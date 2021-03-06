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
	r.HandleFunc(
		"/v2/service_instances/{instance_id}/last_operation",
		pollLastOperation(operations),
	).Methods("GET")

	// provision
	r.HandleFunc("/v2/service_instances/{instance_id}", provision(operations)).Methods("PUT")

	// update service instance
	r.HandleFunc(
		"/v2/service_instances/{instance_id}",
		updateInstance(operations),
	).Methods("PATCH")

	// bind
	r.HandleFunc(
		"/v2/service_instances/{instance_id}/service_bindings/{binding_id}",
		bind(operations),
	).Methods("PUT")

	hostStr := fmt.Sprintf(":%d", port)
	http.ListenAndServe(hostStr, r)
	return nil
}

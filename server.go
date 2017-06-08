package osbserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Run(port int, operations Operations) error {
	r := mux.NewRouter()

	r.HandleFunc("/v2/catalog", getCatalog(operations)).Methods("GET")

	hostStr := fmt.Sprintf(":%d", port)
	http.ListenAndServe(hostStr, r)
	return nil
}

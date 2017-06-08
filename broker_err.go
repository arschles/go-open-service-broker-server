package osbserver

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type brokerErrBody struct {
	Description string `json:"description"`
}

func writeBrokerErr(w http.ResponseWriter, b *BrokerError) {
	w.WriteHeader(b.Code)
	json.NewEncoder(w).Encode(brokerErrBody{Description: b.Description})
}

type BrokerError struct {
	Code        int
	Description string
}

func (b BrokerError) Error() string {
	return fmt.Sprintf("Broker error code %d (%s)", b.Code, b.Description)
}

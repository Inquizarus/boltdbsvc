package handlers

import (
	"encoding/json"
	"net/http"
)

func writeResponse(w http.ResponseWriter, response Response) error {
	bbytes, _ := json.Marshal(response)
	w.Write(bbytes)
	return nil
}

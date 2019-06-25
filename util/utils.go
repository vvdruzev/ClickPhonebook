package util

import (
	"net/http"
	"encoding/json"
)

func ResponseOk(w http.ResponseWriter, body []byte) {


	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func ResponseError(w http.ResponseWriter, code int, message string) {
	body := map[string]string{
		"error": message,
	}
	b,_:= json.Marshal(body)

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
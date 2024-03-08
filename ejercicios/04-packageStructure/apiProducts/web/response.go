package web

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, code int, data any) {
	//set content type
	w.Header().Set("Content-Type", "application/json")

	//set status code
	w.WriteHeader(code)

	//encode the data
	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}
}

package web

import (
	"encoding/json"
	"net/http"
)

func RequestJSON(w http.ResponseWriter, r *http.Request, data any) error {
	//decode the data
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return err
	}

	return nil
}

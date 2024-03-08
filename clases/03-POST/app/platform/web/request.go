package web

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrInvalidContentType = errors.New("invalid content type")
)

func RequestJSON(r *http.Request, ptr any) (err error) {
	//check the content type.
	if r.Header.Get("Content-Type") != "application/json" {
		err = ErrInvalidContentType
		return
	}

	//decode the request body into the ptr.
	if err = json.NewDecoder(r.Body).Decode(ptr); err != nil {
		return
	}

	return
}

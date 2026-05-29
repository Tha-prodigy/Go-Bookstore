package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// parse the request body from json into Go struct
func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal(body, x); err == nil {
			return

		}

	} 
}

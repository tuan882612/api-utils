package apiutils

import (
	"encoding/json"
	"net/http"
)

func Jsonify(w http.ResponseWriter, value interface{}) error {
	w.Header().Set("Content-Type", "application/json")

	if value == nil {
		return nil
	}

	newEncoder := json.NewEncoder(w)

	if err := newEncoder.Encode(value); err != nil {
		return err
	}

	return nil
}

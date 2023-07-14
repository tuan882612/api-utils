package apiutils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Jsonify(t *testing.T) {
	w := httptest.NewRecorder()

	value := map[string]interface{}{
		"key": "value",
	}

	if err := jsonify(w, value); err != nil {
		t.Errorf("Jsonify returned an error: %v", err)
	}

	res := w.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Jsonify returned an invalid status code: %v", res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Jsonify returned an invalid Content-Type: %v", contentType)
	}

	var responseBody map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&responseBody); err != nil {
		t.Errorf("Failed to decode response body: %v", err)
	}

	if len(responseBody) != 1 || responseBody["key"] != "value" {
		t.Errorf("Jsonify returned an invalid response body: %v", responseBody)
	}
}

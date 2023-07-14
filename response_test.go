package apiutils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_WrapRes(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"test": "foo",
		}

		body := NewRes(http.StatusOK, "", data)
		WrapRes(w, body)
	}

	req := httptest.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()
	handler(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("WrapRes returned an invalid status code: %v", res.Code)
	}

	contentType := res.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("WrapRes returned an invalid Content-Type: %v", contentType)
	}

	resBody := &Response{}
	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		t.Errorf("Failed to decode response body: %v", err)
	}

	if res.Body == nil || resBody.Body.(map[string]interface{})["test"] != "foo" {
		t.Errorf("WrapRes returned an invalid response body: %v", resBody)
	}
}

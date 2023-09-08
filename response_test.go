package apiutils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_SendRes(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"test": "foo",
		}

		resp := NewRes(http.StatusOK, "", data)
		resp.SendRes(w)
	}

	req := httptest.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()
	handler(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("SendRes returned an invalid status code: %v", res.Code)
	}

	contentType := res.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("SendRes returned an invalid Content-Type: %v", contentType)
	}

	resBody := &Response{}
	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		t.Errorf("Failed to decode response body: %v", err)
	}

	if res.Body == nil || resBody.Body.(map[string]interface{})["test"] != "foo" {
		t.Errorf("SendRes returned an invalid response body: %v", resBody)
	}
}

func Test_AddHeaders(t *testing.T) {
	cases := []struct {
		name   string
		header map[string]string
		expect []string
	}{
		{
			name: "single header",
			header: map[string]string{
				"X-Test-Header": "foo",
			},
			expect: []string{"X-Test-Header: foo"},
		},
		{
			name: "multiple headers",
			header: map[string]string{
				"X-Test-Header":  "foo",
				"X-Test-Header2": "bar",
			},
			expect: []string{"X-Test-Header: foo", "X-Test-Header2: bar"},
		},
		{
			name: "empty header",
			header: map[string]string{
				"": "foo",
			},
			expect: []string{},
		},
	}

	for _, c := range cases {
		handler := func(w http.ResponseWriter, r *http.Request) {
			resp := NewRes(http.StatusOK, "", nil)
			resp.AddHeader(w, c.header)
			resp.SendRes(w)
		}

		req := httptest.NewRequest("GET", "/", nil)
		res := httptest.NewRecorder()
		handler(res, req)

		if res.Code != http.StatusOK {
			t.Errorf("AddHeader returned an invalid status code: %v", res.Code)
		}

		contentType := res.Header().Get("Content-Type")
		if contentType != "application/json" {
			t.Errorf("AddHeader returned an invalid Content-Type: %v", contentType)
		}

		for _, e := range c.expect {
			if res.Header().Get(e) != c.header[e] {
				t.Errorf("AddHeader returned an invalid header: %v", e)
			}
		}
	}

}

package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/myazahq/go-test-exercise/section3/exercise1/handlers"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorldHandler(t *testing.T) {

	testCases := []struct {
		name         string
		httpMethod   string
		responseCode int
		urlPath      string
	}{
		{
			name:         "success case",
			httpMethod:   "GET",
			responseCode: http.StatusOK,
		},
		{
			name:         "put method not allowed case",
			httpMethod:   "PUT",
			responseCode: http.StatusMethodNotAllowed,
		},
		{
			name:         "post method not allowed case",
			httpMethod:   "POST",
			responseCode: http.StatusMethodNotAllowed,
		},
		{
			name:         "patch method not allowed case",
			httpMethod:   "PATCH",
			responseCode: http.StatusMethodNotAllowed,
		},
		{
			name:         "delete method not allowed case",
			httpMethod:   "DELETE",
			responseCode: http.StatusMethodNotAllowed,
		},
		{
			name:         "url not found case",
			httpMethod:   "GET",
			responseCode: http.StatusNotFound,
			urlPath:      "invalid",
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlers.HelloWorld)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req := httptest.NewRequest(tc.httpMethod, "http://localhost:8080/"+tc.urlPath, nil)

			mux.ServeHTTP(rr, req)

			assert.Equal(t, tc.responseCode, rr.Code)
			if rr.Code == http.StatusOK {
				expected := "Hello, World!\n"
				body, err := io.ReadAll(rr.Body)
				if err != nil {
					t.Fatal(err)
				}
				if string(body) != expected {
					t.Errorf("handler returned unexpected body: got %v want %v", string(body), expected)
				}
			}

		})
	}
}

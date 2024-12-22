package application

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AtariOverlord09/gowebcalc/config"
	"go.uber.org/zap"
)

func TestCalcHandler(t *testing.T) {
	logger, _ := zap.NewProduction()
	app := New(&config.Config{Host: "localhost", Port: 8080}, logger)

	tests := []struct {
		name           string
		method         string
		body           string
		expectedStatus int
		expectedBody   interface{}
		wantError      bool
	}{
		{
			name:           "Valid Expression",
			method:         http.MethodPost,
			body:           `{"expression": "3+4"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   Response{Result: "7.000000"},
			wantError:      false,
		},
		{
			name:           "Invalid Expression",
			method:         http.MethodPost,
			body:           `{"expression": "3+4*("}`,
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   ErrParse,
			wantError:      true,
		},
		{
			name:           "Invalid Method",
			method:         http.MethodGet,
			body:           "",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   ErrMethodNotAllowed,
			wantError:      true,
		},
		{
			name:           "Empty Body",
			method:         http.MethodPost,
			body:           "",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   ErrEmptyBody,
			wantError:      true,
		},
		{
			name:           "Invalid JSON",
			method:         http.MethodPost,
			body:           `{"expression":`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   ErrBadRequest,
			wantError:      true,
		},
		{
			name:           "Division by Zero",
			method:         http.MethodPost,
			body:           `{"expression": "3/0"}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   ErrDivisionByZero,
			wantError:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/", bytes.NewBufferString(tt.body))
			rr := httptest.NewRecorder()

			app.CalcHandler(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			var responseBody interface{}
			if tt.wantError {
				responseBody = &ErrorResponse{}
			} else {
				responseBody = &Response{}
			}

			err := json.NewDecoder(rr.Body).Decode(responseBody)
			if err != nil {
				t.Errorf("failed to decode response body: %v", err)
			}

			if !compareResponses(responseBody, tt.expectedBody) {
				t.Errorf("handler returned unexpected body: got %v want %v", responseBody, tt.expectedBody)
			}
		})
	}
}

func compareResponses(got, want interface{}) bool {
	gotBytes, err := json.Marshal(got)
	if err != nil {
		return false
	}
	wantBytes, err := json.Marshal(want)
	if err != nil {
		return false
	}
	return bytes.Equal(gotBytes, wantBytes)
}

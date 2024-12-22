// I think it is worth separating the errors of the main functionality
// of the calculation and web errors from the API.
// That's why here we have implemented
// a structure that allows to compare errors
// of the calculator and its web representation
package application

import (
	"encoding/json"
	"net/http"

	err "github.com/AtariOverlord09/gowebcalc/pkg/calculation"
)

var (
	ErrMethodNotAllowed = NewErrorResponse("method not allowed", http.StatusMethodNotAllowed)
	ErrBadRequest       = NewErrorResponse("bad request", http.StatusBadRequest)
	ErrInternalServer   = NewErrorResponse("internal server error", http.StatusInternalServerError)
	ErrEmptyBody        = NewErrorResponse("empty body", http.StatusBadRequest)
	ErrDivisionByZero   = NewErrorResponse("division by zero error", http.StatusBadRequest)
	ErrExpression       = NewErrorResponse("expression error", http.StatusUnprocessableEntity)
	ErrParse            = NewErrorResponse("parse error", http.StatusUnprocessableEntity)
	ErrInvalidOperation = NewErrorResponse("invalid operation error", http.StatusUnprocessableEntity)
)

type ErrorResponse struct {
	Msg        string `json:"error"`
	statusCode int
}

func (e ErrorResponse) WriteTo(w http.ResponseWriter) {
	w.WriteHeader(e.statusCode)
	json.NewEncoder(w).Encode(e)
}

func NewErrorResponse(msg string, statusCode int) ErrorResponse {
	return ErrorResponse{
		Msg:        msg,
		statusCode: statusCode,
	}
}

var apiErrors = map[error]ErrorResponse{
	err.ErrParse:            ErrParse,
	err.ErrExpression:       ErrExpression,
	err.ErrInvalidOperation: ErrInvalidOperation,
	err.ErrDivisionByZero:   ErrDivisionByZero,
}

package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/AtariOverlord09/gowebcalc/config"
	"github.com/AtariOverlord09/gowebcalc/internal/middleware"
	"github.com/AtariOverlord09/gowebcalc/pkg/calculation"

	"go.uber.org/zap"
)

type Application struct {
	config *config.Config
	log    *zap.Logger
}

func New(cfg *config.Config, logger *zap.Logger) *Application {
	return &Application{
		config: cfg,
		log:    logger,
	}
}

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result"`
}

func (a *Application) CalcHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		ErrMethodNotAllowed.WriteTo(w)
		return
	}

	defer r.Body.Close()
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		if err.Error() == "EOF" {
			ErrEmptyBody.WriteTo(w)
			return
		}
		ErrBadRequest.WriteTo(w)
		return
	}

	result, err := calculation.Calc(strings.ReplaceAll(req.Expression, " ", ""))
	if err != nil {
		errorResponse, ok := apiErrors[err]
		if ok {
			errorResponse.WriteTo(w)
		} else {
			ErrInternalServer.WriteTo(w)
		}
	} else {
		response := Response{Result: fmt.Sprintf("%f", result)}
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			ErrInternalServer.WriteTo(w)
		}
	}
}

func (a *Application) RunServer() error {
	a.log.Info("Starting server", zap.String("host", a.config.Host), zap.Int("port", a.config.Port))

	http.HandleFunc("/", middleware.Logging(a.CalcHandler, a.log))
	return http.ListenAndServe(a.config.Host+":"+strconv.Itoa(a.config.Port), nil)
}

func (a *Application) Close() {
	a.log.Info("Closing server")
}

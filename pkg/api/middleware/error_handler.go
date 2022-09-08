package middleware

import (
	"net/http"

	errorHandler "github.com/chocogem/slackbot-golang/pkg/api/error"
	"github.com/chocogem/slackbot-golang/pkg/driver/log"
	"github.com/gin-gonic/gin"
)

const (
	InternalServerErrorCode string = "ERR_001"
	AuthenticationErrorCode string = "ERR_002"

	InternalServerErrorMsg string = "Internal Server Error"
	AuthenticationErrorMsg string = "Authentication Error"
)

type ErrorResponse struct {
	Code    string    `json:"code"`
	Message string    `json:"message"`
	Errors  []Message `json:"errors"`
}

type Message struct {
	Message string `json:"message"`
}

type ErrorHandler struct {
	log log.Logger
}

func NewErrorHandler(log log.Logger) *ErrorHandler {
	return &ErrorHandler{log}
}

func (e *ErrorHandler) Handler() gin.HandlerFunc {
	return func(gctx *gin.Context) {
		gctx.Next()
		if len(gctx.Errors) > 0 {
			var response ErrorResponse
			gerr := gctx.Errors[0].Unwrap()
			e.log.Error(gerr.Error())
			switch e := gerr.(type) {
			case *errorHandler.ErrorAuthentication:
				response.Code = AuthenticationErrorCode
				response.Message = AuthenticationErrorMsg
				response.Errors = append(response.Errors, Message{Message: e.Error()})
				gctx.JSON(http.StatusBadRequest, response)
				return
			default:
				response.Code = InternalServerErrorCode
				response.Message = InternalServerErrorMsg
				response.Errors = append(response.Errors, Message{Message: e.Error()})
				gctx.JSON(http.StatusInternalServerError, response)
				return
			}
		}
	}
}

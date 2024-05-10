package server

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	. "value-survey/pkg/errors"
)

const (
	codeSuccess        = 0
	codeServerRunError = iota + 10000
	codeQueryParseError
	codeJsonParseError
)

type HttpServer struct {
	address string
	e       *gin.Engine
}

func NewHttpServer() *HttpServer {
	e := gin.Default()
	return &HttpServer{e: e}
}

func (h *HttpServer) Run(address string) Error {
	if err := NewWithError(codeServerRunError, h.e.Run(address)); err != nil {
		return err
	}
	return nil
}

type Handler[Req, Resp interface{}] func(req *Req) (*Resp, Error)

type RetResult struct {
	Result  interface{}
	Code    int
	Message string
}

func GET[Req, Resp interface{}](server *HttpServer, path string, handler Handler[Req, Resp]) {
	server.e.GET(path, func(context *gin.Context) {
		var req Req
		if err := NewWithError(codeQueryParseError, context.ShouldBind(&req)); err != nil {
			log.Error(err.StackTrace())
			context.JSON(http.StatusBadRequest, &RetResult{
				Code:    codeQueryParseError,
				Message: "query error",
			})
			return
		}

		if err := validInterface(&req); err != nil {
			log.Error(err.StackTrace())
			context.JSON(http.StatusBadRequest, &RetResult{
				Code:    err.Code(),
				Message: err.Message(),
			})
			return
		}

		resp, err := handler(&req)
		if err != nil {
			log.Error(err.StackTrace())
			context.JSON(http.StatusInternalServerError, &RetResult{
				Result:  resp,
				Code:    err.Code(),
				Message: err.Message(),
			})
			return
		}

		context.JSON(http.StatusOK, &RetResult{
			Result:  resp,
			Code:    codeSuccess,
			Message: "success",
		})
	})
}

func POST[Req, Resp interface{}](server *HttpServer, path string, handler Handler[Req, Resp]) {
	server.e.POST(path, func(context *gin.Context) {
		var req Req
		if err := NewWithError(codeJsonParseError, context.ShouldBindJSON(&req)); err != nil {
			log.Error(err.StackTrace())
			context.JSON(http.StatusBadRequest, &RetResult{
				Code:    codeJsonParseError,
				Message: "json parse error",
			})
			return
		}

		if err := validInterface(&req); err != nil {
			log.Error(err.StackTrace())
			context.JSON(http.StatusBadRequest, &RetResult{
				Code:    err.Code(),
				Message: err.Message(),
			})
			return
		}

		resp, err := handler(&req)
		if err != nil {
			log.Error(err.StackTrace())
			context.JSON(http.StatusInternalServerError, &RetResult{
				Result:  resp,
				Code:    err.Code(),
				Message: err.Message(),
			})
			return
		}

		context.JSON(http.StatusOK, &RetResult{
			Result:  resp,
			Code:    codeSuccess,
			Message: "success",
		})
	})
}

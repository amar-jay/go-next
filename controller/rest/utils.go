package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// type of Response object
type Response struct {
	Status int    `json:"status"`
	Msg    string `json:"message"`
	Data   any    `json:"data"`
}

// send an http response
func HttpResponse(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(
		code, Response{
			Status: code,
			Msg:    msg,
			Data:   data,
		},
	)
}

func handleErr(ctx *gin.Context, err error, e string) {
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			HttpResponse(ctx, http.StatusNotFound, e, nil)
			return
		}

		HttpResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
}

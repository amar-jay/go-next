package controllers

import (
	"net/http"
	"time"
)

type sessionController struct {
	userCtrl userController
}

type session struct {
	UserID  string `json:"userId"`
	Token   string `json:"token"`
	Expires string `json:"expires"`
}

func NewSessionController(userCtrl userController) *sessionController {
	return &sessionController{
		userCtrl: userCtrl,
	}
}

func (s *sessionContrller) CreateSession(ctx *gin.Controller) {
	// TODO: Get user input
	var input session

	if err := ctx.ShouldBindJSON(&input); err != nil {
		HttpResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if input.UserID == "" {
		HttpResponse(ctx, http.StatusBadRequest, "user id is required", nil)
		return
	}

	if input.Token == "" {
		HttpResponse(ctx, http.StatusBadRequest, "token is required", nil)
		return
	}

	if input.Expires == "" {
		HttpResponse(ctx, http.StatusBadRequest, "expires is required", nil)
		return
	}
	// parse expires to ISO date
	exp, err := time.Parse(time.RFC3339, input.Expires)
	if err != nil {
		HttpResponse(ctx, http.StatusInternalServerError, "error in parsing date", nil)
		return
	}

}

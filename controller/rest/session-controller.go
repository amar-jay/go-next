package controllers

import (
	"net/http"
	"time"

	models "github.com/amar-jay/go-api-boilerplate/database/domain/session"
	"github.com/gin-gonic/gin"
)

type sessionController struct {
	userCtrl userController
}

type session struct {
	UserID  string `json:"userId"`
	Token   string `json:"sessionToken"`
	Expires string `json:"expires"`
}

func NewSessionController(userCtrl userController) *sessionController {
	return &sessionController{
		userCtrl: userCtrl,
	}
}

func (s *sessionController) CreateSession(ctx *gin.Context) {
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

	if err = s.userCtrl.us.CreateSession(&models.Session{
		UserID:  input.UserID,
		Token:   input.Token,
		Expires: exp,
	}); err != nil {
		HttpResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	input.Expires = exp.String()
	HttpResponse(ctx, http.StatusOK, "session created successfully", input)
}

func (s *sessionController) GetSession(ctx *gin.Context) {
	var token string = ctx.Param("token")
	if token != "" {
		if sess, err := s.userCtrl.us.GetSession(token); err == nil {
			HttpResponse(ctx, http.StatusOK, "session found", session{
				UserID:  sess.UserID,
				Token:   token,
				Expires: time.Now().Add(time.Hour * 1).String(), // 1 hour
			})
			return
		}

		HttpResponse(ctx, http.StatusNotFound, "error fetching session from db", nil)
		return
	}

	HttpResponse(ctx, http.StatusBadRequest, "invalid params", nil)
}

func (s *sessionController) UpdateSession(ctx *gin.Context) {
	// get session from body
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

	if out, err := s.userCtrl.us.UpdateSession(&models.Session{
		UserID:  input.UserID,
		Token:   input.Token,
		Expires: time.Now().Add(time.Hour * 1), // 1 hour
	}); err == nil {
		HttpResponse(ctx, http.StatusOK, "session updated", out)
		return
	}

	HttpResponse(ctx, http.StatusInternalServerError, "error updating session", nil)
	// TODO: Update session in database

}
func (s *sessionController) DeleteSession(ctx *gin.Context) {
	var token string = ctx.Param("token")
	if token != "" {
		// TODO: Delete session from database
		if err := s.userCtrl.us.DeleteSession(token); err != nil {
			HttpResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		HttpResponse(ctx, http.StatusOK, "session deleted", nil)
		return
	}

	HttpResponse(ctx, http.StatusBadRequest, "invalid params", nil)

}

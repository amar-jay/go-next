package controllers

import (
	"net/http"
	"time"

	"github.com/amar-jay/go-api-boilerplate/database/domain/account"
	models "github.com/amar-jay/go-api-boilerplate/database/domain/session"
	"github.com/gin-gonic/gin"
)

type session struct {
	UserID  string `json:"userId"`
	Token   string `json:"sessionToken"`
	Expires string `json:"expires"`
}

// account is the account input type for linking accounts,
// it contains refreshToken, accessToken and stuff
type acc struct {
	ID           string `json:"id"`
	SID          string `json:"sid"`
	ProviderID   string `json:"providerId"`
	Type         string `json:"providerType"`
	AccID        string `json:"providerAccountId"`
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
	Expires      string `json:"accessTokenExpires"`
}

// account input type for unlinking accounts
// it just contains the provider type and the account id
type acc_input struct {
	Type  string `json:"providerType"`
	AccID string `json:"providerAccountId"`
}

func (s *userController) CreateSession(ctx *gin.Context) {
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

	if err = s.us.CreateSession(&models.Session{
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

func (s *userController) GetSession(ctx *gin.Context) {
	var token string = ctx.Query("token")
	if token != "" {
		if sess, err := s.us.GetSession(token); err == nil {
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

func (s *userController) UpdateSession(ctx *gin.Context) {
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

	if out, err := s.us.UpdateSession(&models.Session{
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
func (s *userController) DeleteSession(ctx *gin.Context) {
	var token string = ctx.Param("token")
	if token != "" {
		// TODO: Delete session from database
		if err := s.us.DeleteSession(token); err != nil {
			HttpResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		HttpResponse(ctx, http.StatusOK, "session deleted", nil)
		return
	}

	HttpResponse(ctx, http.StatusBadRequest, "invalid params", nil)

}

func (s *userController) LinkAccount(ctx *gin.Context) {
	// get account from body
	var input acc
	if err := ctx.ShouldBindJSON(&input); err != nil {
		HttpResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if input == (acc{}) || input.Type == "" || input.AccID == "" {
		HttpResponse(ctx, http.StatusBadRequest, "invalid params", nil)
		return
	}

	if err := s.us.LinkAccount(&account.Account{
		ID:                 input.ID,
		SID:                input.SID,
		ProviderId:         input.ProviderID,
		ProviderType:       input.Type,
		AccountID:          input.AccID,
		RefreshToken:       input.RefreshToken,
		AccessToken:        input.AccessToken,
		AccessTokenExpires: input.Expires,
	}); err != nil {
		HttpResponse(ctx, http.StatusInternalServerError, "Database input error: "+err.Error(), nil)
		return
	}

	HttpResponse(ctx, http.StatusOK, "account linked", nil)
}

func (s *userController) UnlinkAccount(ctx *gin.Context) {

	// get account from body
	var input acc_input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		HttpResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if input == (acc_input{}) || input.Type == "" || input.AccID == "" {
		HttpResponse(ctx, http.StatusBadRequest, "invalid params", nil)
		return
	}

	if err := s.us.UnlinkAccount(input.Type, input.AccID); err != nil {
		HttpResponse(ctx, http.StatusInternalServerError, "Database input error: "+err.Error(), nil)
		return
	}
	HttpResponse(ctx, http.StatusOK, "account unlinked", nil)
}

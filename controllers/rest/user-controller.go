package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/amar-jay/go-api-boilerplate/database/domain/user"
	"github.com/amar-jay/go-api-boilerplate/services/authservice"
	"github.com/amar-jay/go-api-boilerplate/services/emailservice"
	"github.com/amar-jay/go-api-boilerplate/services/userservice"
	"github.com/gin-gonic/gin"
)

/**
*  --- Types ---
 */
type LoginUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserInput struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}
type UserOutput struct {
	ID        uint   `json:"ID"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Active    bool   `json:"acive"`
}

type UserUpdateInput struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}
type ErrorOutput struct {
	Msg     string `json:"message"`
	Summary string `json:"summary"`
	Code    int    `json:"statusCode"`
}

/**
* ---- Input Types -----
 */

type UserController interface {
	Register(*gin.Context)
	Update(*gin.Context)
	login(ctx *gin.Context, u *user.User, message string) error
	Login(ctx *gin.Context)
	GetUserByID(*gin.Context)
	GetUsers(*gin.Context)
	GetProfile(*gin.Context)
	ResetPassword(*gin.Context)
	ForgotPassword(*gin.Context)
}

type userController struct {
	us userservice.UserService
	as authservice.AuthService
	es emailservice.EmailService
}

// NewUserService creates a an instance of User Service
func NewUserController(us userservice.UserService, as authservice.AuthService, es emailservice.EmailService) UserController {
	return &userController{us, as, es}
}

/**
* ----- Routes -----
 */

func (userctrl *userController) Login(ctx *gin.Context) {

	// TODO: Get user input
	var input LoginUserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		HttpResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	u := userctrl.login_input_to_User(input)
	// TODO: Get User from Database
	user, err := userctrl.us.Login(&u)
	if err != nil {
		HttpResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// TODO: Check Password
	err = userctrl.us.ComparePassword(input.Password, user.Password)
	if err != nil {
		HttpResponse(ctx, http.StatusBadRequest, "password does not match", nil)
		return
	}
	// TODO: Login
	if err := userctrl.login(ctx, user, "user logged in successfully"); err != nil {
		HttpResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
}

func (userctrl *userController) Register(ctx *gin.Context) {
	//  read the user input
	var userInput RegisterUserInput
	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		HttpResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	u := userctrl.register_input_to_User(userInput)
	// create a user
	if err := userctrl.us.Register(&u); err != nil {
		HttpResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// TODO: send a welcome message by mail
	if err := userctrl.es.Welcome(u.Email); err != nil {
		HttpResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	//  login the user
	if err := userctrl.login(ctx, &u, "Sign up successfully"); err != nil {
		HttpResponse(ctx, http.StatusMethodNotAllowed, err.Error(), nil)
		return
	}
}

func (userctrl *userController) Update(ctx *gin.Context) {
	// read the user id
	input, exists := ctx.Get("user_id")
	if !exists {
		HttpResponse(ctx, http.StatusBadRequest, "Invalid user ID entered", nil)
		return
	}

	// get the user from the database
	user, err := userctrl.us.GetUserByID(input.(uint))
	if err != nil {
		HttpResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	// Read user input
	var userInput UserUpdateInput
	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		HttpResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Check if user is true
	if user.ID != input {
		HttpResponse(ctx, http.StatusUnauthorized, "User Unauthorized", nil)
		return
	}

	//  Update the user Record
	user.FirstName = userInput.FirstName
	user.LastName = userInput.FirstName
	user.Email = userInput.Email
	if err := userctrl.us.Update(user); err != nil {
		HttpResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// Reponse
	userOutput := userctrl.mapToUserOutput(user)
	HttpResponse(ctx, http.StatusAccepted, "ok", userOutput)
}
func (user *userController) ResetPassword(ctx *gin.Context) {
	fmt.Println("🔎 Reset Password controller not implemented")
}

func (user *userController) ForgotPassword(ctx *gin.Context) {
	fmt.Println("🔎 Forgot Password controller not implemented")
}

func (userctrl *userController) GetProfile(ctx *gin.Context) {
	id, exists := ctx.Get("user_id")

	if !exists {
		HttpResponse(ctx, http.StatusBadRequest, "Invalid User ID fetched", nil)
		return
	}

	user, err := userctrl.us.GetUserByID(id.(uint))
	if err != nil {
		HttpResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	userOutput := userctrl.mapToUserOutput(user)
	HttpResponse(ctx, http.StatusOK, "ok", userOutput)
}

func (userctrl *userController) GetUsers(ctx *gin.Context) {
	var usersOut []*UserOutput
	users, err := userctrl.us.GetUsers()
	// map each user to usersOut
	for _, user := range users {
		out := userctrl.mapToUserOutput(user)
		usersOut = append(usersOut, out)
	}

	if err != nil {
		HttpResponse(ctx, http.StatusNotFound, err.Error(), nil)
		return
	}

	HttpResponse(ctx, http.StatusOK, "ok", usersOut)
}

func (userctrl *userController) GetUserByID(ctx *gin.Context) {
	id, err := userctrl.getUserID(ctx.Param("id"))

	if err != nil {
		HttpResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	user, err := userctrl.us.GetUserByID(id)
	if err != nil {
		e := err.Error()
		if strings.Contains(e, "not found") {
			HttpResponse(ctx, http.StatusNotFound, e, nil)
			return
		}
		HttpResponse(ctx, http.StatusNotFound, e, nil)
		return
	}

	userOutput := userctrl.mapToUserOutput(user)
	HttpResponse(ctx, http.StatusOK, "ok", userOutput)

}

/**
* --- Other methods
 */

// converts user input to user model
func (userctrl *userController) login_input_to_User(input LoginUserInput) user.User {

	return user.User{
		Email:    input.Email,
		Password: input.Password,
	}
}

// converts user input to user model
func (userctrl *userController) register_input_to_User(input RegisterUserInput) user.User {

	return user.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
		Role:      input.Role,
		Active:    true,
	}
}

func (userctrl *userController) mapToUserOutput(input *user.User) *UserOutput {
	return &UserOutput{
		ID:        input.ID,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Active:    input.Active,
		Role:      input.Role,
	}
}
func (userctrl *userController) login(ctx *gin.Context, u *user.User, message string) error {
	token, err := userctrl.as.IssueToken(*u)
	if err != nil {
		return err
	}
	//  userOutput := userctrl.mapToUserOutput(u)
	out := gin.H{"token": token}
	HttpResponse(ctx, http.StatusOK, message, out)
	return nil
}

// Get user by id using ID param
func (userctrl *userController) getUserID(IDparam string) (uint, error) {

	userID, err := strconv.Atoi(IDparam)
	if err != nil {
		return 0, errors.New("user id should be a number")
	}

	return uint(userID), nil
}

package user

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/leungyauming/api/app"
	"github.com/leungyauming/api/common/responses"
	"github.com/leungyauming/api/common/security"
	"net/http"
	"net/mail"
)

type RegisterController struct {
	deps *app.Deps
}

func NewRegisterController(deps *app.Deps) *RegisterController {
	return &RegisterController{deps: deps}
}

func (c *RegisterController) Any(ctx *gin.Context) {
	switch ctx.Request.Method {
	case http.MethodPost:
		usernameForm := ctx.PostForm("username")
		emailForm := ctx.PostForm("email")
		passwordForm := ctx.PostForm("password")
		if usernameForm == "" || emailForm == "" || passwordForm == "" {
			responses.SendErrorResponse(ctx,
				http.StatusBadRequest,
				errors.New("username, email and password form value must not be empty"))
			return
		}

		emailAddr, err := mail.ParseAddress(emailForm)
		if err != nil {
			responses.SendErrorResponse(ctx,
				http.StatusBadRequest,
				errors.New("malformed email address"))
			return
		}

		email := emailAddr.Address

		reqCtx := ctx.Request.Context()

		var userAlreadyExists bool
		err = c.deps.Database.QueryRow(reqCtx,
			"SELECT EXISTS(SELECT 1 FROM users WHERE user_username=$1);",
			usernameForm).Scan(&userAlreadyExists)
		if errors.Is(err, context.Canceled) {
			return
		}
		if err != nil {
			responses.SendErrorResponse(ctx,
				http.StatusInternalServerError,
				errors.New("failed to query info from database"))
			return
		}
		if userAlreadyExists {
			responses.SendErrorResponse(ctx,
				http.StatusBadRequest,
				errors.New("user already exists"))
			return
		}

		salt, pwHash, err := security.Argon2idHashPassword(passwordForm, security.DefaultArgon2idParams)
		pwHashEncoded := security.Argon2idEncodePasswordHash(salt, pwHash, security.DefaultArgon2idParams)

		_, err = c.deps.Database.Exec(reqCtx,
			"INSERT INTO users(user_username, user_email, user_password_hash_encoded) VALUES ($1, $2, $3);",
			usernameForm, email, pwHashEncoded)
		if errors.Is(err, context.Canceled) {
			return
		}
		if err != nil {
			responses.SendErrorResponse(ctx,
				http.StatusInternalServerError,
				errors.New("failed to write info to database"))
			return
		}

		responses.SendMsgResponse(ctx, http.StatusOK, "user registered successfully")

	default:
		responses.SendMethodNotAllowed(ctx)
	}
}

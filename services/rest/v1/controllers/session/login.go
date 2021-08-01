package session

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unionofblackbean/api/app"
	"github.com/unionofblackbean/api/common/responses"
	"github.com/unionofblackbean/api/common/security"
	"net/http"
)

type LoginController struct {
	deps *app.Deps
}

func NewLoginController(deps *app.Deps) *LoginController {
	return &LoginController{deps: deps}
}

func (c *LoginController) Any(ctx *gin.Context) {
	switch ctx.Request.Method {
	case http.MethodPost:
		reqUsername := ctx.PostForm("username")
		reqPassword := ctx.PostForm("password")
		if reqUsername == "" || reqPassword == "" {
			responses.SendErrorResponse(ctx,
				http.StatusUnauthorized,
				errors.New("username and password form value must not be empty"))
			return
		}

		reqCtx := ctx.Request.Context()

		var dbPwHashEncoded string
		err := c.deps.Postgres.QueryRow(reqCtx,
			"SELECT user_password_hash_encoded FROM users WHERE user_username=$1;",
			reqUsername,
		).Scan(&dbPwHashEncoded)
		if errors.Is(err, context.Canceled) {
			return
		}
		if err != nil {
			responses.SendErrorResponse(ctx,
				http.StatusInternalServerError,
				errors.New("failed to query info from database"))
			return
		}

		passwordIsCorrect, err := security.Argon2idVerifyPassword(reqPassword, dbPwHashEncoded)
		if err != nil {
			responses.SendErrorResponse(ctx,
				http.StatusInternalServerError,
				fmt.Errorf("failed to verify password -> %v", err))
			return
		}
		if !passwordIsCorrect {
			responses.SendErrorResponse(ctx,
				http.StatusUnauthorized,
				errors.New("incorrect credentials"))
			return
		}

		sessionIDBytes, err := security.GenerateRandomBytes(32)
		if err != nil {
			responses.SendErrorResponse(ctx,
				http.StatusInternalServerError,
				fmt.Errorf("failed to generate session id -> %v", err))
			return
		}
		sessionIDString := base64.RawURLEncoding.EncodeToString(sessionIDBytes)

		_, err = c.deps.Postgres.Exec(reqCtx,
			"INSERT INTO sessions(session_id, session_ip, user_username) VALUES ($1, $2, $3);",
			sessionIDString, ctx.ClientIP(), reqUsername)
		if errors.Is(err, context.Canceled) {
			return
		}
		if err != nil {
			responses.SendErrorResponse(ctx,
				http.StatusInternalServerError,
				errors.New("failed to save session info"))
			return
		}

		responses.SendJsonResponse(ctx,
			http.StatusOK,
			"success",
			&gin.H{
				"session_id": sessionIDString,
			})

	default:
		responses.SendMethodNotAllowed(ctx)
	}
}

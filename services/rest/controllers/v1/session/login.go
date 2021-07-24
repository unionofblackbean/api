package session

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/leungyauming/api/app"
	"github.com/leungyauming/api/common/responses"
	"github.com/leungyauming/api/common/security"
	"net/http"
)

type LoginController struct {
	deps *app.Deps
}

func NewLoginController(deps *app.Deps) *LoginController {
	c := new(LoginController)
	c.deps = deps

	return c
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

		var dbPwHashEncoded string
		err := c.deps.Database.QueryRow(context.Background(),
			"SELECT user_password_hash_encoded FROM users WHERE user_username=$1;",
			reqUsername,
		).Scan(&dbPwHashEncoded)
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

		sessionID, err := uuid.NewRandom()
		if err != nil {
			responses.SendErrorResponse(ctx,
				http.StatusInternalServerError,
				fmt.Errorf("failed to generate session id -> %v", err))
			return
		}

		sessionIDString := sessionID.String()

		_, err = c.deps.Database.Exec(context.Background(),
			"INSERT INTO sessions(session_id, session_ip, user_username) VALUES ($1, $2, $3);",
			sessionIDString, ctx.ClientIP(), reqUsername)
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

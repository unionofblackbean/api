package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/leungyauming/api/common/responses"
	"github.com/leungyauming/api/common/security"
	"net/http"
)

type LoginController struct {
	db *pgxpool.Pool
}

func NewLoginController(db *pgxpool.Pool) *LoginController {
	c := new(LoginController)
	c.db = db

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
		err := c.db.QueryRow(context.Background(),
			"SELECT password_hash_encoded FROM users WHERE username=$1;",
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

		_, err = c.db.Exec(context.Background(),
			"INSERT INTO sessions(id, username, ip) VALUES ($1, $2, $3);",
			sessionIDString, reqUsername, ctx.ClientIP())
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

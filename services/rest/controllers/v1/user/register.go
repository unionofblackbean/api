package user

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/leungyauming/api/common/responses"
	"github.com/leungyauming/api/common/security"
	"net/http"
	"net/mail"
)

type RegisterController struct {
	db *pgxpool.Pool
}

func NewRegisterController(db *pgxpool.Pool) *RegisterController {
	c := new(RegisterController)
	c.db = db

	return c
}

func (c *RegisterController) Post(ctx *gin.Context) {
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

	var userAlreadyExists bool
	err = c.db.QueryRow(context.Background(),
		"SELECT EXISTS(SELECT 1 FROM users WHERE username=$1);",
		usernameForm).Scan(&userAlreadyExists)
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

	_, err = c.db.Exec(context.Background(),
		"INSERT INTO users(username, email, password_hash_encoded) VALUES ($1, $2, $3);",
		usernameForm, email, pwHashEncoded)
	if err != nil {
		responses.SendErrorResponse(ctx,
			http.StatusInternalServerError,
			errors.New("failed to write info to database"))
		return
	}

	responses.SendMsgResponse(ctx, http.StatusOK, "user registered successfully")
}

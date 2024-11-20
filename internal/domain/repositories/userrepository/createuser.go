package userrepository

import (
	"context"
	"log/slog"

	"github.com/d3cie/pubnode/internal/domain/errs"
	"github.com/d3cie/pubnode/internal/domain/models"
	repository "github.com/d3cie/pubnode/internal/domain/repositories"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserInput struct {
	Username     string
	Email        string
	AuthProvider string

	Name      *string
	GithubID  *string
	AvatarUrl *string
	Password  *string
}

func (r *userRepository) CreateUser(ctx context.Context, input CreateUserInput, uow *repository.UowTx) (*models.User, errs.Error) {
	r.logger.InfoContext(ctx, "CreateUser: creating new user")
	r.logger.DebugContext(ctx, "CreateUser: input", slog.Any("input", input))

	db := r.db
	if uow != nil {
		db = uow.DB
	}

	user := models.User{
		Username:      input.Username,
		Email:         input.Email,
		Name:          input.Name,
		AvatarUrl:     input.AvatarUrl,
		GithubID:      input.GithubID,
		AuthProviders: []string{input.AuthProvider},
	}

	if input.Password != nil {
		var err error
		user.PasswordHash, err = bcrypt.GenerateFromPassword([]byte(*input.Password), 10)
		if err != nil {
			r.logger.ErrorContext(ctx, "CreateUser: error generating password hash", slog.Any("error", err))
			return nil, errs.ErrInternal
		}
	}

	err := db.Create(&user).Error
	if err != nil {
		if db.IsErrUniqueConstraintViolation(err, []string{"email"}) {
			return nil, errs.ErrUserWithEmailAlreadyExists
		}
		if db.IsErrUniqueConstraintViolation(err, []string{"github_id"}) {
			return nil, errs.ErrUserWithEmailAlreadyExists
		}
		r.logger.ErrorContext(ctx, "CreateUser: error creating user", slog.Any("error", err))
		return nil, errs.ErrInternal
	}

	r.logger.DebugContext(ctx, "CreateUser: user", slog.Any("user", user))
	return &user, nil
}

package authservice

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/d3cie/pubnode/internal/domain/errs"
	"github.com/d3cie/pubnode/internal/domain/models"
	"github.com/d3cie/pubnode/internal/domain/repositories/sessionrepository"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Email    *string
	Password *string

	GithubID     *string
	AuthProvider string

	UserAgent *string
	IpAddress *string
}

func (s *authService) Login(ctx context.Context, input LoginInput) (*models.Session, errs.Error) {
	s.logger.InfoContext(ctx, "Login: Loging in user")
	s.logger.DebugContext(ctx, "Login: input", slog.Any("input", input))

	var user *models.User
	var session *models.Session
	var err errs.Error

	if input.AuthProvider == "email" {
		user, err = s.userRepository.GetUserByEmail(ctx, *input.Email)
		if err != nil {
			if errors.Is(err, errs.ErrUserNotFound) {
				return nil, errs.ErrInvalidCredentials
			}
			return nil, err
		}
		if bcryptErr := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(*input.Password)); bcryptErr != nil {
			if errors.Is(bcryptErr, bcrypt.ErrMismatchedHashAndPassword) {
				return nil, errs.ErrInvalidCredentials
			}
			s.logger.ErrorContext(ctx, "Login: error comparing password hash", slog.Any("error", bcryptErr))
			return nil, errs.ErrInternal
		}
	}

	if input.AuthProvider == "github" {
		panic("not implemented")
	}

	session, err = s.sessionRepository.CreateSession(ctx, sessionrepository.CreateSessionInput{
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(SESSION_EXPIRES_IN),
		UserAgent: input.UserAgent,
		IpAddress: input.IpAddress,
	}, nil)
	if err != nil {
		return nil, err
	}

	s.logger.InfoContext(ctx, "Login: session created", slog.Any("session", session), slog.Any("user", user))
	return session, nil
}

package authservice

import (
	"context"
	"log/slog"
	"time"

	"github.com/d3cie/pubnode/internal/domain/errs"
	"github.com/d3cie/pubnode/internal/domain/models"
	"github.com/d3cie/pubnode/internal/domain/repositories/sessionrepository"
	"github.com/d3cie/pubnode/internal/domain/repositories/userrepository"

	repository "github.com/d3cie/pubnode/internal/domain/repositories"
)

type RegisterInput struct {
	Username     string
	Email        string
	AuthProvider string

	Password  *string
	Name      *string
	AvatarUrl *string
	GithubID  *string
	UserAgent *string
	IpAddress *string
}

type RegisterOutput struct {
	User    models.User
	Session models.Session
}

func (s *authService) Register(ctx context.Context, input RegisterInput) (*RegisterOutput, errs.Error) {
	s.logger.InfoContext(ctx, "Register: signing up user")
	s.logger.DebugContext(ctx, "Register: input", slog.Any("input", input))

	uow := s.uowProvider.New()

	var user *models.User
	var session *models.Session
	var err errs.Error

	err = uow.Do(ctx, func(tx *repository.UowTx) errs.Error {
		user, err = s.userRepository.CreateUser(ctx, userrepository.CreateUserInput{
			Email:        input.Email,
			Username:     input.Username,
			Name:         input.Name,
			Password:     input.Password,
			GithubID:     input.GithubID,
			AuthProvider: input.AuthProvider,
			AvatarUrl:    input.AvatarUrl,
		}, tx)
		if err != nil {
			return err
		}
		session, err = s.sessionRepository.CreateSession(ctx, sessionrepository.CreateSessionInput{
			UserID:    user.ID,
			ExpiresAt: time.Now().Add(SESSION_EXPIRES_IN),
			UserAgent: input.UserAgent,
			IpAddress: input.IpAddress,
		}, tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	s.logger.DebugContext(ctx, "Register: user, session", slog.Any("user", user), slog.Any("session", session))
	return &RegisterOutput{
		User:    *user,
		Session: *session,
	}, nil
}

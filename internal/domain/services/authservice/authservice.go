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

const (
	SESSION_EXPIRES_IN    = time.Hour * 24 * 30
	REFRESH_SESSION_AFTER = time.Hour * 24 * 15
)

type AuthService interface {
	Login(ctx context.Context, input LoginInput) (*models.Session, errs.Error)
	Register(ctx context.Context, input RegisterInput) (*RegisterOutput, errs.Error)
}

type authService struct {
	uowProvider repository.UOWProvider

	userRepository    userrepository.UserRepository
	sessionRepository sessionrepository.SessionRepository

	logger *slog.Logger
}

func New(
	uowProvider repository.UOWProvider,

	userRepository userrepository.UserRepository,
	sessionRepository sessionrepository.SessionRepository,

	logger *slog.Logger,
) AuthService {
	return &authService{
		uowProvider: uowProvider,

		userRepository:    userRepository,
		sessionRepository: sessionRepository,

		logger: logger.With("service", "authservice"),
	}
}

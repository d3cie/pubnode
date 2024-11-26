package userrepository

import (
	"context"
	"log/slog"

	"github.com/d3cie/pubnode/internal/domain/errs"
	"github.com/d3cie/pubnode/internal/domain/models"
	repository "github.com/d3cie/pubnode/internal/domain/repositories"
	"github.com/d3cie/pubnode/internal/infra/db"
)

type UserRepository interface {
	GetUser(ctx context.Context, id string) (*models.User, errs.Error)
	GetUserByGithubID(ctx context.Context, id string) (*models.User, errs.Error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, errs.Error)
	CreateUser(ctx context.Context, input CreateUserInput, uow *repository.UowTx) (*models.User, errs.Error)
}

type userRepository struct {
	logger *slog.Logger
	db     *db.DB
}

func New(db *db.DB, logger *slog.Logger) UserRepository {
	return &userRepository{
		db:     db,
		logger: logger.With("repository", "userrepository"),
	}
}

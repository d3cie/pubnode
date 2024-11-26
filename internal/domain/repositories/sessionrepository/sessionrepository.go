package sessionrepository

import (
	"context"
	"log/slog"

	"github.com/d3cie/pubnode/internal/domain/errs"
	"github.com/d3cie/pubnode/internal/domain/models"
	repository "github.com/d3cie/pubnode/internal/domain/repositories"
	"github.com/d3cie/pubnode/internal/infra/db"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, input CreateSessionInput, uow *repository.UowTx) (*models.Session, errs.Error)
	GetSession(sessionID string) (*models.Session, errs.Error)
}

type sessionRepository struct {
	logger *slog.Logger
	db     *db.DB
}

func New(db *db.DB, logger *slog.Logger) SessionRepository {
	return &sessionRepository{
		db:     db,
		logger: logger.With("repository", "sessionrepository"),
	}
}

package sessionrepository

import (
	"context"
	"log/slog"
	"time"

	"github.com/d3cie/pubnode/internal/domain/errs"
	"github.com/d3cie/pubnode/internal/domain/models"
	repository "github.com/d3cie/pubnode/internal/domain/repositories"
)

type CreateSessionInput struct {
	UserID    string
	ExpiresAt time.Time
	IpAddress *string
	UserAgent *string
}

func (r *sessionRepository) CreateSession(ctx context.Context, input CreateSessionInput, uow *repository.UowTx) (*models.Session, errs.Error) {
	r.logger.InfoContext(ctx, "CreateSession: creating new session")
	r.logger.DebugContext(ctx, "CreateSession: input", slog.Any("input", input))

	db := r.db
	if uow != nil {
		db = uow.DB
	}

	session := models.Session{
		UserID:    input.UserID,
		ExpiresAt: input.ExpiresAt,
		IpAddress: input.IpAddress,
		UserAgent: input.UserAgent,
	}

	err := db.Create(&session).Error
	if err != nil {
		r.logger.ErrorContext(ctx, "CreateSession: error creating session", slog.Any("error", err))
		return nil, errs.ErrInternal
	}

	r.logger.DebugContext(ctx, "CreateSession: session", slog.Any("session", session))
	return &session, nil
}

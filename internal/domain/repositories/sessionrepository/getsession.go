package sessionrepository

import (
	"log/slog"

	"github.com/d3cie/pubnode/internal/domain/errs"
	"github.com/d3cie/pubnode/internal/domain/models"
)

func (r *sessionRepository) GetSession(sessionID string) (*models.Session, errs.Error) {
	r.logger.Info("GetSession: getting session")
	r.logger.Debug("GetSession: sessionID", slog.String("sessionID", sessionID))

	session := &models.Session{}
	err := r.db.Where("session_id = ?", sessionID).First(session).Error
	if err != nil {
		return nil, errs.ErrInternal
	}

	r.logger.Debug("GetSession: session", slog.Any("session", session))
	return session, nil
}

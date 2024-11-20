package userrepository

import (
	"context"
	"log/slog"

	"github.com/d3cie/pubnode/internal/domain/errs"
	"github.com/d3cie/pubnode/internal/domain/models"
)

func (r *userRepository) GetUser(ctx context.Context, id string) (*models.User, errs.Error) {
	r.logger.Info("GetUser: getting user")
	r.logger.Debug("GetUser: id", slog.String("id", id))

	user := &models.User{}
	err := r.db.Where("id = ?", id).First(user).Error
	if err != nil {
		if r.db.IsErrNotFound(err) {
			return nil, errs.ErrUserNotFound
		}
		r.logger.Error("GetUser: an error occured fetching user", slog.Any("error", err))
		return nil, errs.ErrInternal
	}

	r.logger.Debug("GetUser: user", slog.Any("user", user))
	return user, nil
}

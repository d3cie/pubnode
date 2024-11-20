package userrepository

import (
	"context"
	"log/slog"

	"github.com/d3cie/pubnode/internal/domain/errs"
	"github.com/d3cie/pubnode/internal/domain/models"
)

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, errs.Error) {
	r.logger.Info("GetUserByEmail: getting user")
	r.logger.Debug("GetUserByEmail: email", slog.String("email", email))

	user := &models.User{}
	err := r.db.Where("email = ?", email).First(user).Error
	if err != nil {
		if r.db.IsErrNotFound(err) {
			return nil, errs.ErrUserNotFound
		}
		r.logger.Error("GetUserByEmail: an error occured fetching user", slog.Any("error", err))
		return nil, errs.ErrInternal
	}

	r.logger.Debug("GetUserByEmail: user", slog.Any("user", user))
	return user, nil
}

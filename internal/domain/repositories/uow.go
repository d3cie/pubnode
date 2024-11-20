package repository

import (
	"context"
	"log/slog"

	"github.com/d3cie/pubnode/internal/domain/errs"
	"github.com/d3cie/pubnode/internal/infra/db"
)

type UOWProvider interface {
	New() UOW
}

type UOW interface {
	Do(ctx context.Context, fn func(*UowTx) errs.Error) errs.Error
}

type uowProvider struct {
	logger *slog.Logger
	db     *db.DB
}

type uow struct {
	logger *slog.Logger
	db     *db.DB
}

type UowTx struct {
	DB *db.DB
}

func NewUOWProvider(db *db.DB, logger *slog.Logger) UOWProvider {
	return &uowProvider{
		db:     db,
		logger: logger.With("manager", "uow"),
	}
}

func (m *uowProvider) New() UOW {
	return &uow{
		db:     m.db,
		logger: m.logger,
	}
}

func (u *uow) Do(ctx context.Context, fn func(tx *UowTx) errs.Error) errs.Error {
	u.logger.InfoContext(ctx, "Do: executing a unit of work")
	db := u.db.Tx()
	xerr := fn(&UowTx{DB: db})
	if xerr != nil {
		db.Rollback()
		return xerr
	} else {
		db.Commit()
		return nil
	}
}

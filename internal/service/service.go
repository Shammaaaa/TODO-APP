package service

import (
	"context"

	"github.com/shamil/todo-app/internal/infrastructure/database"
	"github.com/shamil/todo-app/internal/infrastructure/database/postgres"
	"github.com/shamil/todo-app/pkg/drop"
)

type Service struct {
	*drop.Impl

	Pool database.Pool
}

func New(ctx context.Context, opt *Options) (*Service, error) {
	s := &Service{}
	s.Impl = drop.NewContext(ctx)

	var err error

	s.Pool, err = postgres.NewPool(s.Context(), opt.Database)
	if err != nil {
		return nil, err
	}
	s.AddDropper(s.Pool.(*postgres.Pool))

	return s, nil
}

package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shamil/todo-app/internal/infrastructure/database"
)

type Pool struct {
	db *pgxpool.Pool
}

// Builder возвращает пул соединений pgx.
func (c *Pool) Builder() *pgxpool.Pool {
	return c.db
}

// Drop закрывает пул соединений.
func (c *Pool) Drop() error {
	c.db.Close()
	return nil
}

// DropMsg возвращает сообщение о закрытии пула.
func (c *Pool) DropMsg() string {
	return "close database: pool closed"
}

// NewPool создает новый пул соединений с использованием pgx.
func NewPool(ctx context.Context, opt *database.Opt) (*Pool, error) {
	// Настройка конфигурации пула
	config, err := pgxpool.ParseConfig(opt.ConnectionString())
	if err != nil {
		return nil, err
	}

	// Установка параметров пула
	config.MaxConns = int32(opt.MaxOpenConns)
	config.MinConns = int32(opt.MaxIdleConns)
	config.MaxConnLifetime = opt.MaxConnMaxLifetime

	// Создание пула
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	// Проверка соединения
	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := pool.Ping(pingCtx); err != nil {
		return nil, err
	}

	return &Pool{db: pool}, nil
}

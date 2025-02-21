package database

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool interface {
	Builder() *pgxpool.Pool
}

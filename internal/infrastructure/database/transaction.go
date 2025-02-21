package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// TxFn — это функция, которая выполняется внутри транзакции.
type TxFn func(tx pgx.Tx) error

// WithTransaction выполняет переданную функцию внутри транзакции.
// Если функция возвращает ошибку, транзакция откатывается.
// Если функция завершается успешно, транзакция коммитится.
func WithTransaction(ctx context.Context, pool *pgxpool.Pool, fn TxFn) (err error) {
	// Начинаем транзакцию
	tx, err := pool.Begin(ctx)
	if err != nil {
		return err
	}

	// Отложенный вызов для обработки отката или коммита
	defer func() {
		if p := recover(); p != nil {
			// В случае паники откатываем транзакцию и повторно вызываем панику
			_ = tx.Rollback(ctx)
			panic(p)
		} else if err != nil {
			// Если есть ошибка, откатываем транзакцию
			_ = tx.Rollback(ctx)
		} else {
			// Если все хорошо, коммитим транзакцию
			err = tx.Commit(ctx)
		}
	}()

	// Выполняем переданную функцию внутри транзакции
	err = fn(tx)
	return err
}

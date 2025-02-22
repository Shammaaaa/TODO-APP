package database

import (
	"context"
	"log"

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
		log.Printf("Failed to begin transaction: %v", err)
		return err
	}

	// Отложенный вызов для обработки отката или коммита
	defer func() {
		if p := recover(); p != nil {
			rollbackErr := tx.Rollback(ctx)
			if rollbackErr != nil {
				log.Printf("Failed to rollback transaction after panic: %v", rollbackErr)
			}
			panic(p)
		} else if err != nil {
			rollbackErr := tx.Rollback(ctx)
			if rollbackErr != nil {
				log.Printf("Failed to rollback transaction: %v", rollbackErr)
			}
		} else {
			err = tx.Commit(ctx)
			if err != nil {
				log.Printf("Failed to commit transaction: %v", err)
			}
		}
	}()

	err = fn(tx)
	if err != nil {
		log.Printf("Transaction function failed: %v", err)
	}
	return err
}

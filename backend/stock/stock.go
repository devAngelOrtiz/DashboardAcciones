package stock

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Stock struct {
	ID         uuid.UUID `json:"id"`
	Ticker     string    `json:"ticker"`
	Company    string    `json:"company"`
	Brokerage  string    `json:"brokerage"`
	Action     string    `json:"action"`
	RatingFrom string    `json:"rating_from"`
	RatingTo   string    `json:"rating_to"`
	TargetFrom string    `json:"target_from"`
	TargetTo   string    `json:"target_to"`
	Time       time.Time `json:"time"`
}

type PaginatedStocks struct {
	Stocks []Stock `json:"stocks"`
	Total  int     `json:"total"`
	Pages  int     `json:"pages"`
	Page   int     `json:"page"`
}

const defaultPage = 1
const defaultTotal = 10

func InsertData(conn *pgx.Conn, stock Stock) error {
	_, err := conn.Exec(context.Background(),
		`INSERT INTO stock (ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to, time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		stock.Ticker, stock.Company, stock.Brokerage, stock.Action, stock.RatingFrom, stock.RatingTo, stock.TargetFrom, stock.TargetTo, stock.Time)

	if err != nil {
		return fmt.Errorf("error inserting stock: %v", err)
	}

	return nil
}

const (
	querySearchStocks = `
		SELECT * FROM stock 
		WHERE brokerage ILIKE $1 
		ORDER BY time DESC 
		LIMIT $2 OFFSET $3`
	queryCountStocksSearch = `SELECT COUNT(*) FROM stock WHERE brokerage ILIKE $1`

	queryStocks = `
		SELECT * FROM stock 
		ORDER BY time DESC 
		LIMIT $1 OFFSET $2`
	queryCountStocks = `SELECT COUNT(*) FROM stock`
)

func GetStocks(ctx context.Context, conn *pgx.Conn, page, total int, search string) (*PaginatedStocks, error) {
	offset := (page - 1) * total

	var count int

	var query string
	var args []interface{}
	var countQuery string
	var countArgs []interface{}

	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("error al iniciar la transacción: %w", err)
	}
	defer tx.Rollback(ctx)

	if search != "" {
		query = querySearchStocks
		countQuery = queryCountStocksSearch
		args = []interface{}{"%" + search + "%", total, offset}
		countArgs = []interface{}{"%" + search + "%"}

	} else {
		query = queryStocks
		countQuery = queryCountStocks
		args = []interface{}{total, offset}
	}

	err = tx.QueryRow(ctx, countQuery, countArgs...).Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("error al obtener el total de registros: %w", err)
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta de stocks: %w", err)
	}

	defer rows.Close()

	var stocks []Stock
	for rows.Next() {
		var s Stock
		err := rows.Scan(&s.ID, &s.Ticker, &s.Company, &s.Brokerage, &s.Action, &s.RatingFrom, &s.RatingTo, &s.TargetFrom, &s.TargetTo, &s.Time)
		if err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %w", err)
		}
		stocks = append(stocks, s)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar sobre las filas: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("error al confirmar la transacción: %w", err)
	}

	pages := 1
	if total > 0 {
		pages = int(math.Ceil(float64(count) / float64(total)))
	}

	return &PaginatedStocks{
		Stocks: stocks,
		Total:  count,
		Pages:  pages,
		Page:   page,
	}, nil
}

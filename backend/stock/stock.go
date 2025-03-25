package stock

import (
	"context"
	"fmt"
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

func InsertData(conn *pgx.Conn, stock Stock) error {
	_, err := conn.Exec(context.Background(),
		`INSERT INTO stock (ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to, time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		stock.Ticker, stock.Company, stock.Brokerage, stock.Action, stock.RatingFrom, stock.RatingTo, stock.TargetFrom, stock.TargetTo, stock.Time)

	if err != nil {
		return fmt.Errorf("error inserting stock: %v", err)
	}

	return nil
}

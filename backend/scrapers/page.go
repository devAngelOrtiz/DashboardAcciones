package scrapers

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func SavePageHistory(conn *pgx.Conn, nextPage string) error {
	insertSQL := `INSERT INTO page (page) VALUES ($1);`

	_, err := conn.Exec(context.Background(), insertSQL, nextPage)
	if err != nil {
		return fmt.Errorf("error inserting page: %v", err)
	}

	return nil
}

func GetLastPage(conn *pgx.Conn) string {
	var lastPage string
	selectSQL := `SELECT page FROM page ORDER BY queryAt DESC LIMIT 1;`

	err := conn.QueryRow(context.Background(), selectSQL).Scan(&lastPage)
	if err != nil {
		return ""
	}

	return lastPage
}

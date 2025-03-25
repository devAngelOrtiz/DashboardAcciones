package common

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func GetDBConnection() (*pgx.Conn, error) {
	dbURL := GetEnv("DATABASE_URL")

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		return nil, fmt.Errorf("error conectando a la base de datos: %v", err)
	}

	createStockTable(conn)
	createPageTable(conn)
	return conn, nil

}

func createStockTable(conn *pgx.Conn) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS stock (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		ticker STRING NOT NULL,
		company STRING NOT NULL,
		brokerage STRING NOT NULL,
		action STRING NOT NULL,
		rating_from STRING NOT NULL,
		rating_to STRING NOT NULL,
		target_from STRING NOT NULL,
		target_to STRING NOT NULL,
		time TIMESTAMPTZ DEFAULT NOW()
	);
	`
	_, err := conn.Exec(context.Background(), createTableSQL)
	if err != nil {
		log.Fatalf("error creating table stock: %v", err)
	}

	return nil
}

func createPageTable(conn *pgx.Conn) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS page (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		page STRING NOT NULL,
		queryAt TIMESTAMPTZ DEFAULT NOW()
	);
	`
	_, err := conn.Exec(context.Background(), createTableSQL)
	if err != nil {
		log.Fatalf("error creating table page: %v", err)
	}

	return nil
}

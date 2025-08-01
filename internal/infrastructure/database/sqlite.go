package infrastructure

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func InitDB(path string) *sql.DB {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		log.Fatal("failed to open db: ", err)
	}

	schema := `
	CREATE TABLE IF NOT EXISTS payments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		order_id TEXT NOT NULL,
		payment_gateway TEXT NOT NULL,
		payment_type TEXT NOT NULL,
		external_id TEXT NOT NULL UNIQUE,
		amount REAL NOT NULL,
		status TEXT CHECK(status IN ('PENDING', 'SUCCESS', 'FAILED', 'EXPIRED')) NOT NULL,
		paid_at DATETIME,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS idempotency_keys (
		key TEXT PRIMARY KEY,
		request_hash TEXT,
		response_body TEXT,
		status_code INTEGER,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal("failed to create table: ", err)
	}

	return db
}

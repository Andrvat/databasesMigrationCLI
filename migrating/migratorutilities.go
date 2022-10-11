package migrating

import (
	"database/sql"
	"log"
)

func CloseConn(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}

func Finally(err error) {
	if err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}
}

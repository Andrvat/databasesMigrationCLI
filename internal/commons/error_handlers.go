package commons

import (
	"database/sql"
	"log"
)

func CloseConn(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}

func HandleMigrationErr(err error, onSuccessMsg string) {
	if err != nil && err.Error() != "no change" {
		log.Fatal(err)
	} else {
		log.Println(onSuccessMsg)
	}
}

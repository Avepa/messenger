package mysql

import (
	"database/sql"
	"time"
)

func UserAdd(db *sql.DB, username string) (id int64, err error) {
	result, err := db.Exec(
		"INSERT INTO `User` (`username`, `created_at`) VALUES (?, ?)",
		username,
		time.Now().UTC(),
	)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return
}

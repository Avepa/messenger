package mysql

import (
	"database/sql"
	"time"
)

type NewMessage struct {
	Chat   int64  `json:"chat"`
	Author int64  `json:"author"`
	Text   string `json:"text"`
}

type Messages struct {
	ID         int64  `json:"id"`
	Chat       int64  `json:"chat"`
	Author     int64  `json:"author"`
	Text       string `json:"text"`
	Created_at string `json:"created_at"`
}

func MessagesAdd(db *sql.DB, message NewMessage) (id int64, err error) {
	result, err := db.Exec(
		"INSERT INTO `Message` (`chat`, `author`, `text`, `created_at`) VALUES (?, ?, ?, ?)",
		message.Chat,
		message.Author,
		message.Text,
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

func MessagesGet(db *sql.DB, chat int64) (messages []Messages, err error) {
	rows, err := db.Query(
		"SELECT `id`, `author`, `text`, `created_at` FROM `Message` WHERE `chat` = ? ORDER BY `created_at` ASC",
		chat,
	)
	if err != nil {
		return []Messages{}, err
	}
	defer rows.Close()

	for rows.Next() {
		message := Messages{Chat: chat}
		if err = rows.Scan(&message.ID, &message.Author, &message.Text, &message.Created_at); err != nil {
			return []Messages{}, err
		}
		messages = append(messages, message)
	}
	return
}

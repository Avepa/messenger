package mysql

import (
	"database/sql"
	"time"
)

type UserChats struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Users      []int64 `json:"users"`
	Created_at string  `json:"created_at"`
}

func ChatsAdd(db *sql.DB, chatName string, users []int64) (id int64, err error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	result, err := tx.Exec(
		"INSERT INTO `Chat` (`name`, `created_at`) VALUES (?, ?)",
		chatName,
		time.Now().UTC(),
	)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	for _, userID := range users {
		_, err = tx.Exec(
			"INSERT INTO `Users` (`Chat_id`, `User_id`) VALUES (?, ?)",
			id,
			userID,
		)
		if err != nil {
			return 0, err
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return
}

func ChatsGet(db *sql.DB, user int64) (chats []UserChats, err error) {
	rows, err := db.Query(
		"SELECT `Chat`.`id`, `Chat`.`name`, `Chat`.`created_at` FROM `Chat`"+
			"    INNER JOIN `Message` ON `Message`.`chat` = `Chat`.`id` "+
			"				INNER JOIN `Users` ON `Users`.`Chat_id` = `Chat`.`id`"+
			"       	WHERE `Users`.`User_id` = ? "+
			"       	GROUP BY `Chat`.`id` "+
			"       	ORDER BY MAX(`Message`.`created_at`) DESC;",
		user,
	)
	if err != nil {
		return []UserChats{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var chat UserChats
		if err = rows.Scan(&chat.ID, &chat.Name, &chat.Created_at); err != nil {
			return []UserChats{}, err
		}

		usersRows, err := db.Query(
			"SELECT `User_id` FROM `Users` WHERE `Chat_id` = ?",
			chat.ID,
		)
		if err != nil {
			return []UserChats{}, err
		}

		for usersRows.Next() {
			var user int64
			if err = usersRows.Scan(&user); err != nil {
				return []UserChats{}, err
			}
			chat.Users = append(chat.Users, user)
		}
		chats = append(chats, chat)
		usersRows.Close()
	}
	return
}

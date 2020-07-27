package url

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Avepa/messenger/src/mysql"
)

type Handler struct {
	DB *sql.DB
}

type ID struct {
	ID int64 `json:"id"`
}

type UserName struct {
	Username string `json:"username"`
}

func (handlers *Handler) UserAdd(w http.ResponseWriter, r *http.Request) {
	var username UserName
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&username); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := mysql.UserAdd(handlers.DB, username.Username)
	if err != nil {
		log.Println(err)
		errorDesc := http.StatusText(http.StatusInternalServerError) + "\n" + err.Error()
		http.Error(w, errorDesc, http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	response := ID{ID: id}
	if err := encoder.Encode(response); err != nil {
		log.Println(err)
		errorDesc := http.StatusText(http.StatusInternalServerError) + "\n" + err.Error()
		http.Error(w, errorDesc, http.StatusInternalServerError)
		return
	}
}

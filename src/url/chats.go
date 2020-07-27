package url

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Avepa/messenger/src/mysql"
)

type UserID struct {
	ID int64 `json:"user"`
}

type NewChat struct {
	ChatName string  `json:"name"`
	Users    []int64 `json:"users"`
}

func (Handlers *Handler) ChatsAdd(w http.ResponseWriter, r *http.Request) {
	var newChat NewChat
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newChat); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := mysql.ChatsAdd(Handlers.DB, newChat.ChatName, newChat.Users)
	if err != nil {
		log.Println(err)
		errorDesc := http.StatusText(http.StatusInternalServerError) + "\n" + err.Error()
		http.Error(w, errorDesc, http.StatusInternalServerError)
		return
	}

	response := ID{ID: id}
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		log.Println(err)
		errorDesc := http.StatusText(http.StatusInternalServerError) + "\n" + err.Error()
		http.Error(w, errorDesc, http.StatusInternalServerError)
		return
	}
}

func (Handlers *Handler) ChatsGet(w http.ResponseWriter, r *http.Request) {
	var userID UserID
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userID); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	chats, err := mysql.ChatsGet(Handlers.DB, userID.ID)
	if err != nil {
		log.Println(err)
		errorDesc := http.StatusText(http.StatusInternalServerError) + "\n" + err.Error()
		http.Error(w, errorDesc, http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(chats); err != nil {
		log.Println(err)
		errorDesc := http.StatusText(http.StatusInternalServerError) + "\n" + err.Error()
		http.Error(w, errorDesc, http.StatusInternalServerError)
		return
	}
}

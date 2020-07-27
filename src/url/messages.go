package url

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Avepa/messenger/src/mysql"
)

type ChatID struct {
	ID int64 `json:"chat"`
}

func (Handlers *Handler) MessagesAdd(w http.ResponseWriter, r *http.Request) {
	var message mysql.NewMessage
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&message); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := mysql.MessagesAdd(Handlers.DB, message)
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

func (Handlers *Handler) MessagesGet(w http.ResponseWriter, r *http.Request) {
	var chat ChatID
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&chat); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	messages, err := mysql.MessagesGet(Handlers.DB, chat.ID)
	if err != nil {
		log.Println(err)
		errorDesc := http.StatusText(http.StatusInternalServerError) + "\n" + err.Error()
		http.Error(w, errorDesc, http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(messages); err != nil {
		log.Println(err)
		errorDesc := http.StatusText(http.StatusInternalServerError) + "\n" + err.Error()
		http.Error(w, errorDesc, http.StatusInternalServerError)
		return
	}
}

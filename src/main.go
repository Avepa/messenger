package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Avepa/messenger/src/url"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	DBHOST := os.Getenv("DATABASE_HOST")
	if DBHOST == "" {
		DBHOST = "127.0.0.1"
	}

	db, err := sql.Open("mysql", "root:12345@tcp("+DBHOST+":3306)/mydb")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(10)
	err = db.Ping()
	if err != nil {
		log.Println(err)
	}

	r := mux.NewRouter()
	handlers := &url.Handler{DB: db}
	r.HandleFunc("/users/add", handlers.UserAdd).Methods("POST")
	r.HandleFunc("/chats/add", handlers.ChatsAdd).Methods("POST")
	r.HandleFunc("/messages/add", handlers.MessagesAdd).Methods("POST")
	r.HandleFunc("/chats/get", handlers.ChatsGet).Methods("POST")
	r.HandleFunc("/messages/get", handlers.MessagesGet).Methods("POST")

	if err := http.ListenAndServe(":9000", r); err != nil {
		log.Fatal(err)
	}
}

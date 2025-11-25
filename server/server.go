package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/ToastedGMS/go-contact-book/controller"
	"github.com/ToastedGMS/go-contact-book/repository"
	_ "github.com/mattn/go-sqlite3"
)

func RunServer() {
	const dbFilePath = "./contacts.db"
	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		log.Fatalf("Error opening SQLite database at %s: %v", dbFilePath, err)
	}
	defer db.Close()
	repo := &repository.SQLRepository{DB: db}
	if err := repo.SetupTables(); err != nil {
		log.Fatalf("Error setting up contacts table: %v", err)
	}
	controller := &controller.Controller{Repo: repo}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			controller.ServerStartHandler(w, r)
			return
		}

		controller.UnknownRouteHandler(w, r)
	})
	mux.HandleFunc("GET /contacts", controller.ListContactsHandler)
	mux.HandleFunc("POST /contacts", controller.AddContactHandler)
	mux.HandleFunc("DELETE /contacts/{ID}", controller.DeleteContactHandler)
	mux.HandleFunc("PATCH /contacts/{ID}", controller.EditContactHandler)

	const port = "8080"
	fmt.Printf("Listening on port %s\n", port)

	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

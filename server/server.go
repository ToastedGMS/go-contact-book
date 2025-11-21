package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ToastedGMS/go-contact-book/controller"
	"github.com/ToastedGMS/go-contact-book/repository"
)

func RunServer() {
	repo := &repository.JSONrepository{FilePath: "contacts.json"}
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

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

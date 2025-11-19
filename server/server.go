package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ToastedGMS/go-contact-book/controller"
)

func RunServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			controller.ServerStartHandler(w, r)
			return
		}

		controller.UnknownRouteHandler(w, r)
	})
	http.HandleFunc("/contacts", controller.ListContactsHandler)
	http.HandleFunc("/contacts/add", controller.AddContactHandler)

	const port = "8080"
	fmt.Printf("Listening on port %s\n", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

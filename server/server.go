package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ToastedGMS/go-contact-book/contactbook"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func RunServer(){
	http.HandleFunc("/", handler)
	http.HandleFunc("/contacts", ListContactsHandler)

	const port = "8080"
	fmt.Printf("Listening on port %s\n", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func ListContactsHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return 
	}

	contacts, err := contactbook.ListContacts()
	if err != nil {
		http.Error(w, "Error retrieving contacts", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(contacts); err != nil {
		log.Printf("Error encoding contacts to JSON: %v", err)
	}
}
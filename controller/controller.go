package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ToastedGMS/go-contact-book/contactbook"
)

func ServerStartHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func UnknownRouteHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This route does not exist", http.StatusNotFound)
}

func ListContactsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Query().Get("name") == "" {

		contacts, err := contactbook.ListContacts()
		if err != nil {
			http.Error(w, "Error retrieving contacts", http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(contacts); err != nil {
			log.Printf("Error encoding contacts to JSON: %v", err)
		}
	} else {
		queryParams := r.URL.Query()
		name := queryParams.Get("name")

		results := contactbook.SearchContacts(name)

		if err := json.NewEncoder(w).Encode(results); err != nil {
			log.Printf("Error encoding search results to JSON: %v", err)
		}
	}
}

func AddContactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var contact contactbook.Contact

	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = contactbook.AddContact(contact.Name, contact.Phone)
	if err != nil {
		http.Error(w, "Error adding contact", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "{\"message\": \"Contact added successfully\"}")
}

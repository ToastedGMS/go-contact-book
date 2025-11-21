package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ToastedGMS/go-contact-book/service"
)

func ServerStartHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func UnknownRouteHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This route does not exist", http.StatusNotFound)
}

func ListContactsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.URL.Query().Get("name") == "" {

		contacts, err := service.ListContacts()
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

		results := service.SearchContacts(name)

		if err := json.NewEncoder(w).Encode(results); err != nil {
			log.Printf("Error encoding search results to JSON: %v", err)
		}
	}
}

func AddContactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var contact service.Contact

	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = service.AddContact(contact.Name, contact.Phone)
	if err != nil {
		http.Error(w, "Error adding contact", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "{\"message\": \"Contact added successfully\"}")
}

func DeleteContactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ID := r.PathValue("ID")
	num, err := strconv.Atoi(ID)
	if err != nil {
		http.Error(w, "Internal Server Error, please try again", http.StatusInternalServerError)
		return
	}

	err = service.DeleteContact(num)
	if err != nil {
		http.Error(w, "Error during contact deletion", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{\"message\": \"Contact deleted successfully\"}")

}

func EditContactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var contact service.Contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	id := r.PathValue("ID")
	num, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Internal Server Error, please try again", http.StatusInternalServerError)
		return
	}

	err = service.EditContact(num, contact.Name, contact.Phone)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{\"message\": \"Contact edited successfully\" }")
}

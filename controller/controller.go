package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ToastedGMS/go-contact-book/contactbook"
)

func ServerStartHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World!")
}

func UnknownRouteHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This route does not exist", http.StatusNotFound)
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
package contactbook

import (
	"encoding/json"
	"fmt"
	"os"
)

type ContactBook struct {
	Contacts []Contact `json:"contacts"`
}

type Contact struct {
	Name string `json:"Name"`
	Phone string `json:"Phone"`
}

var contacts = []Contact{
	{Name: "Hulk", Phone: "777-777-7777"},
	{Name: "Arana", Phone: "131-313-1313"},
}

func AddContact(name, phone string) error{
	contact := Contact{Name: name, Phone: phone}
	existingContacts, err := ListContacts()
	if err != nil {
		return err
	}
	updatedContacts := append(existingContacts, contact)

	contactsData := ContactBook{Contacts: updatedContacts}
	byteData, err := json.Marshal(contactsData)
	if err != nil {
		return err
	}

	err = os.WriteFile("contacts.json", byteData, 0666)

	if err != nil {
		return err
	}

	fmt.Println("Contact added successfully.")
	return nil
} 

func ListContacts() ([]Contact, error){
	fmt.Println("Accessing Contact Book...")
	jsonData, err := os.ReadFile("contacts.json")
	if err != nil {
		return nil, err
	}
	var contactBook ContactBook
	err = json.Unmarshal(jsonData, &contactBook)
	if err != nil {
		return nil, err
	}
	return contactBook.Contacts, nil

}

func SearchContacts(name string) []Contact{
	results, err := ListContacts()
	if err != nil {
		fmt.Println("Error reading contacts")
		return nil
	}
	var filteredResults []Contact
	for _, contact := range results {
		if contact.Name == name {
			filteredResults = append(filteredResults, contact)
		}
	}
	return filteredResults
}


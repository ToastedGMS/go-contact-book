package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type ContactBook struct {
	Contacts []Contact `json:"contacts"`
}

type Contact struct {
	ID    int    `json:"ID"`
	Name  string `json:"Name"`
	Phone string `json:"Phone"`
}

func nextID(contacts []Contact) int {

	maxID := 0
	for _, c := range contacts {
		if c.ID > maxID {
			maxID = c.ID
		}
	}

	return maxID + 1
}

func checkExistingContact(name string, contacts []Contact) bool {
	for _, c := range contacts {
		if strings.EqualFold(name, c.Name) {
			return true
		}
	}
	return false
}

func AddContact(name, phone string) error {

	fileBytes, err := os.ReadFile("contacts.json")
	if err != nil {
		return err
	}

	var contactBook ContactBook
	err = json.Unmarshal(fileBytes, &contactBook)
	if err != nil {
		return err
	}

	existingContacts := contactBook.Contacts
	if checkExistingContact(name, existingContacts) {
		return errors.New("Contact already exists")
	} else {
		contact := Contact{ID: nextID(existingContacts), Name: name, Phone: phone}
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
}

func ListContacts() ([]Contact, error) {
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

func SearchContacts(name string) []Contact {
	results, err := ListContacts()
	if err != nil {
		fmt.Println("Error reading contacts")
		return nil
	}
	var filteredResults []Contact
	for _, contact := range results {
		if strings.EqualFold(contact.Name, name) {
			filteredResults = append(filteredResults, contact)
		}
	}
	return filteredResults
}

func DeleteContact(ID int) error {
	fileBytes, err := os.ReadFile("contacts.json")
	if err != nil {
		return err
	}

	var contactBook ContactBook
	err = json.Unmarshal(fileBytes, &contactBook)
	if err != nil {
		return err
	}

	updatedContacts := []Contact{}
	for _, contact := range contactBook.Contacts {
		if contact.ID != ID {
			updatedContacts = append(updatedContacts, contact)
		}
	}
	contactsData := ContactBook{Contacts: updatedContacts}
	byteData, err := json.Marshal(contactsData)
	if err != nil {
		return err
	}

	err = os.WriteFile("contacts.json", byteData, 0666)

	if err != nil {
		return err
	} else {
		fmt.Println("Contact deleted successfully.")
		return nil
	}
}

func EditContact(ID int, name string, phone string) error {
	fileBytes, err := os.ReadFile("contacts.json")
	if err != nil {
		return err
	}

	var contactBook ContactBook
	err = json.Unmarshal(fileBytes, &contactBook)
	if err != nil {
		return err
	}

	found := false
	for i, contact := range contactBook.Contacts {
		if contact.ID == ID {
			contactBook.Contacts[i].Name = name
			contactBook.Contacts[i].Phone = phone
			found = true
			break
		}
	}
	if !found {
		return errors.New("Contact not found")
	}

	byteData, err := json.Marshal(contactBook)
	if err != nil {
		return err
	}

	err = os.WriteFile("contacts.json", byteData, 0666)
	if err != nil {
		return err
	}

	return nil
}

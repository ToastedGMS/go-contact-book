package service

import (
	"errors"
	"strings"

	"github.com/ToastedGMS/go-contact-book/models"
	"github.com/ToastedGMS/go-contact-book/repository"
)

func nextID(contacts []models.Contact) int {

	maxID := 0
	for _, c := range contacts {
		if c.ID > maxID {
			maxID = c.ID
		}
	}

	return maxID + 1
}

func checkExistingContact(name string, contacts []models.Contact) bool {
	for _, c := range contacts {
		if strings.EqualFold(name, c.Name) {
			return true
		}
	}
	return false
}

func appendContact(name, phone string, existingContacts []models.Contact) []models.Contact {

	contact := models.Contact{ID: nextID(existingContacts), Name: name, Phone: phone}
	updatedContacts := append(existingContacts, contact)
	return updatedContacts

}

func AddContact(name, phone string, repo repository.Repository) error {

	contactBook, err := repo.Read()
	if err != nil {
		return err
	}
	existingContacts := contactBook.Contacts

	if checkExistingContact(name, existingContacts) {

		return errors.New("contact already exists")

	} else {

		contactBook.Contacts = appendContact(name, phone, existingContacts)

		err = repo.Write(contactBook)
		if err != nil {
			return err
		}

		return nil
	}
}

func ListContacts(repo repository.Repository) ([]models.Contact, error) {

	contactBook, err := repo.Read()
	if err != nil {
		return nil, err
	}

	return contactBook.Contacts, nil

}

func SearchContacts(name string, repo repository.Repository) []models.Contact {
	results, err := ListContacts(repo)
	if err != nil {
		return nil
	}
	var filteredResults []models.Contact
	for _, contact := range results {
		if strings.EqualFold(contact.Name, name) {
			filteredResults = append(filteredResults, contact)
		}
	}
	return filteredResults
}

func DeleteContact(ID int, repo repository.Repository) error {

	contactBook, err := repo.Read()
	if err != nil {
		return err
	}

	updatedContacts := []models.Contact{}
	for _, contact := range contactBook.Contacts {
		if contact.ID != ID {
			updatedContacts = append(updatedContacts, contact)
		}
	}

	contactsData := models.ContactBook{Contacts: updatedContacts}
	err = repo.Write(contactsData)
	if err != nil {
		return err
	}

	return nil

}

func EditContact(ID int, name string, phone string, repo repository.Repository) error {

	contactBook, err := repo.Read()
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
		return errors.New("contact not found")
	}

	err = repo.Write(contactBook)
	if err != nil {
		return err
	}

	return nil
}

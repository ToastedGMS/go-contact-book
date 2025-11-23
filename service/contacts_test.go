package service

import (
	"errors"
	"reflect"
	"testing"

	"github.com/ToastedGMS/go-contact-book/models"
	"github.com/ToastedGMS/go-contact-book/repository"
)

// Define local copies of the unexported errors returned by the service functions
var errContactExists = errors.New("contact already exists")
var errContactNotFound = errors.New("contact not found")

func TestListContactsSuccess(t *testing.T) {

	expectedContacts := []models.Contact{{ID: 1, Name: "John Doe", Phone: "67"}}

	var mock repository.MockRepository
	mock.Data = models.ContactBook{Contacts: expectedContacts}

	contacts, err := ListContacts(&mock)
	if err != nil {
		t.Fatalf("ListContacts failed. Expected success, got: %v", err)
	}

	if !reflect.DeepEqual(contacts, expectedContacts) {
		t.Errorf("Returned contacts do not match expected data.")
	}

}

func TestListContactsFailure(t *testing.T) {

	expectedError := errors.New("Fatal error running ListContacts")

	var mock repository.MockRepository
	mock.ReadError = expectedError

	contacts, err := ListContacts(&mock)
	if err == nil {
		t.Errorf("Expected error from ListContacts, but received nil")
	}

	if contacts != nil {
		t.Errorf("Expected ListContacts to return nil, received %v", contacts)
	}

	if !errors.Is(err, expectedError) {
		t.Errorf("Error mismatch. Expected %v, got %v", expectedError, err)
	}
}

func TestAddContactSuccess(t *testing.T) {
	initialContacts := []models.Contact{{ID: 1, Name: "John Doe", Phone: "123"}}
	newName := "Jane Doe"
	newPhone := "456"

	var mock repository.MockRepository
	mock.Data = models.ContactBook{Contacts: initialContacts}

	err := AddContact(newName, newPhone, &mock)
	if err != nil {
		t.Fatalf("AddContact failed. Expected success, got: %v", err)
	}

	// Check if the contact was added and written to the mock repository
	expectedContacts := []models.Contact{
		{ID: 1, Name: "John Doe", Phone: "123"},
		{ID: 2, Name: newName, Phone: newPhone}, // New ID should be 2
	}

	if !reflect.DeepEqual(mock.Data.Contacts, expectedContacts) {
		t.Errorf("Contact book data mismatch after AddContact. Expected:\n%v\nGot:\n%v", expectedContacts, mock.Data.Contacts)
	}
}

func TestAddContactReadFailure(t *testing.T) {
	expectedError := errors.New("read error for AddContact")
	var mock repository.MockRepository
	mock.ReadError = expectedError

	err := AddContact("New Contact", "999", &mock)
	if err == nil {
		t.Errorf("Expected error from AddContact due to Read failure, but received nil")
	}

	if !errors.Is(err, expectedError) {
		t.Errorf("Error mismatch. Expected %v, got %v", expectedError, err)
	}
}

func TestAddContactWriteFailure(t *testing.T) {
	expectedError := errors.New("write error for AddContact")
	var mock repository.MockRepository
	mock.Data = models.ContactBook{Contacts: []models.Contact{}}
	mock.WriteError = expectedError

	err := AddContact("New Contact", "999", &mock)
	if err == nil {
		t.Errorf("Expected error from AddContact due to Write failure, but received nil")
	}

	if !errors.Is(err, expectedError) {
		t.Errorf("Error mismatch. Expected %v, got %v", expectedError, err)
	}
}

func TestAddContactAlreadyExists(t *testing.T) {
	existingName := "Jane Doe"
	initialContacts := []models.Contact{
		{ID: 1, Name: "John Doe", Phone: "123"},
		{ID: 2, Name: existingName, Phone: "456"},
	}

	var mock repository.MockRepository
	mock.Data = models.ContactBook{Contacts: initialContacts}

	err := AddContact(existingName, "789", &mock) // Try to add existing name
	if err == nil {
		t.Errorf("Expected 'contact already exists' error from AddContact, but received nil")
	}

	expectedMsg := errContactExists.Error()
	if err.Error() != expectedMsg {
		t.Errorf("Error message mismatch. Expected '%s', got '%s'", expectedMsg, err.Error())
	}

	// Check that the data was not modified
	if !reflect.DeepEqual(mock.Data.Contacts, initialContacts) {
		t.Errorf("Contact book data modified when adding existing contact. Expected:\n%v\nGot:\n%v", initialContacts, mock.Data.Contacts)
	}
}

func TestSearchContactsFound(t *testing.T) {
	searchName := "John Doe"
	initialContacts := []models.Contact{
		{ID: 1, Name: searchName, Phone: "123"},
		{ID: 2, Name: "Jane Doe", Phone: "456"},
		{ID: 3, Name: "john doe", Phone: "789"}, // Case insensitive match
	}

	expectedResults := []models.Contact{
		{ID: 1, Name: searchName, Phone: "123"},
		{ID: 3, Name: "john doe", Phone: "789"},
	}

	var mock repository.MockRepository
	mock.Data = models.ContactBook{Contacts: initialContacts}

	results := SearchContacts(searchName, &mock)
	if !reflect.DeepEqual(results, expectedResults) {
		t.Errorf("SearchContacts returned unexpected results. Expected:\n%v\nGot:\n%v", expectedResults, results)
	}
}

func TestSearchContactsNotFound(t *testing.T) {
	searchName := "Alice"
	initialContacts := []models.Contact{
		{ID: 1, Name: "John Doe", Phone: "123"},
		{ID: 2, Name: "Jane Doe", Phone: "456"},
	}

	var mock repository.MockRepository
	mock.Data = models.ContactBook{Contacts: initialContacts}

	results := SearchContacts(searchName, &mock)
	if len(results) != 0 {
		t.Errorf("SearchContacts should return empty slice for not found contact, got: %v", results)
	}
}

func TestSearchContactsReadFailure(t *testing.T) {
	expectedError := errors.New("read error for SearchContacts")
	var mock repository.MockRepository
	mock.ReadError = expectedError

	results := SearchContacts("Test", &mock)

	if results != nil {
		t.Errorf("Expected SearchContacts to return nil on Read failure, got: %v", results)
	}
	// Note: The SearchContacts function handles the error internally by returning nil
	// because it calls ListContacts which returns nil and the error, and SearchContacts
	// checks if the error is non-nil and returns nil.
}

func TestDeleteContactSuccess(t *testing.T) {
	IDToDelete := 2
	initialContacts := []models.Contact{
		{ID: 1, Name: "John Doe", Phone: "123"},
		{ID: IDToDelete, Name: "Jane Doe", Phone: "456"},
		{ID: 3, Name: "Peter Pan", Phone: "789"},
	}

	expectedContacts := []models.Contact{
		{ID: 1, Name: "John Doe", Phone: "123"},
		{ID: 3, Name: "Peter Pan", Phone: "789"},
	}

	var mock repository.MockRepository
	mock.Data = models.ContactBook{Contacts: initialContacts}

	err := DeleteContact(IDToDelete, &mock)
	if err != nil {
		t.Fatalf("DeleteContact failed. Expected success, got: %v", err)
	}

	// Check if the contact was deleted and written to the mock repository
	if !reflect.DeepEqual(mock.Data.Contacts, expectedContacts) {
		t.Errorf("Contact book data mismatch after DeleteContact. Expected:\n%v\nGot:\n%v", expectedContacts, mock.Data.Contacts)
	}
}

func TestDeleteContactNotFound(t *testing.T) {
	IDToDelete := 99
	initialContacts := []models.Contact{
		{ID: 1, Name: "John Doe", Phone: "123"},
		{ID: 2, Name: "Jane Doe", Phone: "456"},
	}

	// Deleting a non-existent contact should still succeed and the list should remain the same
	expectedContacts := initialContacts

	var mock repository.MockRepository
	mock.Data = models.ContactBook{Contacts: initialContacts}

	err := DeleteContact(IDToDelete, &mock)
	if err != nil {
		t.Fatalf("DeleteContact failed (for non-existent ID). Expected success, got: %v", err)
	}

	// Check that the data was not modified (or rather, the result is the same as initial)
	if !reflect.DeepEqual(mock.Data.Contacts, expectedContacts) {
		t.Errorf("Contact book data modified unexpectedly after DeleteContact (for non-existent ID). Expected:\n%v\nGot:\n%v", expectedContacts, mock.Data.Contacts)
	}
}

func TestDeleteContactReadFailure(t *testing.T) {
	expectedError := errors.New("read error for DeleteContact")
	var mock repository.MockRepository
	mock.ReadError = expectedError

	err := DeleteContact(1, &mock)
	if err == nil {
		t.Errorf("Expected error from DeleteContact due to Read failure, but received nil")
	}

	if !errors.Is(err, expectedError) {
		t.Errorf("Error mismatch. Expected %v, got %v", expectedError, err)
	}
}

func TestDeleteContactWriteFailure(t *testing.T) {
	expectedError := errors.New("write error for DeleteContact")
	var mock repository.MockRepository
	mock.Data = models.ContactBook{Contacts: []models.Contact{{ID: 1, Name: "Test", Phone: "1"}}}
	mock.WriteError = expectedError

	err := DeleteContact(1, &mock)
	if err == nil {
		t.Errorf("Expected error from DeleteContact due to Write failure, but received nil")
	}

	if !errors.Is(err, expectedError) {
		t.Errorf("Error mismatch. Expected %v, got %v", expectedError, err)
	}
}

func TestEditContactSuccess(t *testing.T) {
	IDToEdit := 2
	newName := "Jane Doe Updated"
	newPhone := "999-000-111"

	initialContacts := []models.Contact{
		{ID: 1, Name: "John Doe", Phone: "123"},
		{ID: IDToEdit, Name: "Jane Doe", Phone: "456"},
		{ID: 3, Name: "Peter Pan", Phone: "789"},
	}

	expectedContacts := []models.Contact{
		{ID: 1, Name: "John Doe", Phone: "123"},
		{ID: IDToEdit, Name: newName, Phone: newPhone},
		{ID: 3, Name: "Peter Pan", Phone: "789"},
	}

	var mock repository.MockRepository
	mock.Data = models.ContactBook{Contacts: initialContacts}

	err := EditContact(IDToEdit, newName, newPhone, &mock)
	if err != nil {
		t.Fatalf("EditContact failed. Expected success, got: %v", err)
	}

	// Check if the contact was edited and written to the mock repository
	if !reflect.DeepEqual(mock.Data.Contacts, expectedContacts) {
		t.Errorf("Contact book data mismatch after EditContact. Expected:\n%v\nGot:\n%v", expectedContacts, mock.Data.Contacts)
	}
}

func TestEditContactNotFound(t *testing.T) {
	IDToEdit := 99
	initialContacts := []models.Contact{
		{ID: 1, Name: "John Doe", Phone: "123"},
		{ID: 2, Name: "Jane Doe", Phone: "456"},
	}

	var mock repository.MockRepository
	mock.Data = models.ContactBook{Contacts: initialContacts}

	err := EditContact(IDToEdit, "Non-existent", "000", &mock)
	if err == nil {
		t.Errorf("Expected 'contact not found' error from EditContact, but received nil")
	}

	expectedMsg := errContactNotFound.Error()
	if err.Error() != expectedMsg {
		t.Errorf("Error message mismatch. Expected '%s', got '%s'", expectedMsg, err.Error())
	}

	// Check that the data was not modified
	if !reflect.DeepEqual(mock.Data.Contacts, initialContacts) {
		t.Errorf("Contact book data modified unexpectedly after EditContact (for non-existent ID). Expected:\n%v\nGot:\n%v", initialContacts, mock.Data.Contacts)
	}
}

func TestEditContactReadFailure(t *testing.T) {
	expectedError := errors.New("read error for EditContact")
	var mock repository.MockRepository
	mock.ReadError = expectedError

	err := EditContact(1, "Test", "1", &mock)
	if err == nil {
		t.Errorf("Expected error from EditContact due to Read failure, but received nil")
	}

	if !errors.Is(err, expectedError) {
		t.Errorf("Error mismatch. Expected %v, got %v", expectedError, err)
	}
}

func TestEditContactWriteFailure(t *testing.T) {
	expectedError := errors.New("write error for EditContact")
	var mock repository.MockRepository
	mock.Data = models.ContactBook{Contacts: []models.Contact{{ID: 1, Name: "Test", Phone: "1"}}}
	mock.WriteError = expectedError

	err := EditContact(1, "Updated", "2", &mock)
	if err == nil {
		t.Errorf("Expected error from EditContact due to Write failure, but received nil")
	}

	if !errors.Is(err, expectedError) {
		t.Errorf("Error mismatch. Expected %v, got %v", expectedError, err)
	}
}

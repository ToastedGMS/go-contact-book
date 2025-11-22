package service

import (
	"errors"
	"reflect"
	"testing"

	"github.com/ToastedGMS/go-contact-book/models"
	"github.com/ToastedGMS/go-contact-book/repository"
)

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

package main

import (
	"slices"
	"testing"
)

func TestAddContact(t *testing.T) {
	contacts = []Contact{} // Reset contacts before test
	AddContact("Hulk", "777-777-7777")

	expected := []Contact{
		{Name: "Hulk", Phone: "777-777-7777"},
	}

	if slices.Equal(contacts, expected) != true {
		t.Errorf("AddContact(\"Hulk\", \"777-777-7777\") failed." )
	}
}

func TestListContacts(t *testing.T){
	contacts = []Contact{
		{Name: "Hulk", Phone: "777-777-7777"},
		{Name: "Arana", Phone: "131-313-1313"},
	}

	var returnedContacts, err = ListContacts()
	if err != nil {
		t.Errorf("ListContacts failed.")
	}

	if slices.Equal(returnedContacts, contacts) != true {
		t.Errorf("ListContacts failed.")
	}
}

func TestSearchContacts(t *testing.T){
	contacts = []Contact{
		{Name: "Hulk", Phone: "777-777-7777"},
		{Name: "Arana", Phone: "131-313-1313"},
	}

	var expected = []Contact{
		{Name: "Hulk", Phone: "777-777-7777"},
	}

	if slices.Equal(SearchContacts("Hulk"), expected) != true {
		t.Errorf("SearchContacts failed.")
	}
}
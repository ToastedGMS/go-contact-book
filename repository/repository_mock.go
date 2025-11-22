package repository

import "github.com/ToastedGMS/go-contact-book/models"

type MockRepository struct {
	Data       models.ContactBook
	ReadError  error
	WriteError error
}

func (m *MockRepository) Read() (models.ContactBook, error) {
	if m.ReadError != nil {
		return models.ContactBook{}, m.ReadError
	}

	return m.Data, nil
}

func (m *MockRepository) Write(data models.ContactBook) error {
	if m.WriteError != nil {
		return m.WriteError
	}

	m.Data = data
	return nil

}

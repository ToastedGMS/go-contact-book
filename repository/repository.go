package repository

import (
	"encoding/json"
	"os"

	"github.com/ToastedGMS/go-contact-book/models"
)

type Repository interface {
	Read() (models.ContactBook, error)
	Write(models.ContactBook) error
}

type JSONrepository struct {
	FilePath string
}

func (receiver *JSONrepository) Read() (models.ContactBook, error) {

	data, err := os.ReadFile(receiver.FilePath)
	if err != nil {
		return models.ContactBook{}, err
	}

	var contactBook models.ContactBook
	err = json.Unmarshal(data, &contactBook)
	if err != nil {
		return models.ContactBook{}, err
	}

	return contactBook, nil
}

func (receiver *JSONrepository) Write(data models.ContactBook) error {

	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(receiver.FilePath, byteData, 0666)
	if err != nil {
		return err
	}

	return nil

}

package repository

import (
	"database/sql"
	"log"

	"github.com/ToastedGMS/go-contact-book/models"
)

type SQLRepository struct {
	DB *sql.DB
}

func (r *SQLRepository) SetupTables() error {

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS contacts (
	id INTEGER NOT NULL PRIMARY KEY,
	name TEXT UNIQUE NOT NULL,
	phone TEXT NOT NULL)
	`

	_, err := r.DB.Exec(sqlStmt)
	if err != nil {
		log.Printf("Error setting up contacts table: %v", err)
	}

	return err

}

func (r *SQLRepository) Read() (models.ContactBook, error) {

	rows, err := r.DB.Query("SELECT id, name, phone FROM contacts ORDER BY id ASC")
	if err != nil {
		return models.ContactBook{}, err
	}
	defer rows.Close()

	var book models.ContactBook
	book.Contacts = []models.Contact{}

	for rows.Next() {
		var contact models.Contact

		err = rows.Scan(&contact.ID, &contact.Name, &contact.Phone)
		if err != nil {
			return models.ContactBook{}, err
		}
		book.Contacts = append(book.Contacts, contact)
	}

	if err = rows.Err(); err != nil {
		return models.ContactBook{}, err
	}

	return book, nil

}

func (r *SQLRepository) Write(data models.ContactBook) error {

	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.Exec("DELETE FROM contacts"); err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO contacts (id, name, phone) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, contact := range data.Contacts {
		_, err = stmt.Exec(contact.ID, contact.Name, contact.Phone)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

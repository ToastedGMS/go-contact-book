package models

type ContactBook struct {
	Contacts []Contact `json:"contacts"`
}

type Contact struct {
	ID    int    `json:"ID"`
	Name  string `json:"Name"`
	Phone string `json:"Phone"`
}

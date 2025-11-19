package main

import (
	"fmt"
	"os"

	"github.com/ToastedGMS/go-contact-book/contactbook"
	"github.com/ToastedGMS/go-contact-book/server"
)

func main(){
	server.RunServer()
	if len(os.Args) < 2 {
		PrintUsage()
	} else if os.Args[1] == "add" && len(os.Args) == 4 {
		contactbook.AddContact(os.Args[2], os.Args[3])
	} else if os.Args[1] == "list" {
		contacts, err := contactbook.ListContacts()
		if err != nil {
			fmt.Println("Error listing contacts")
		}
		for _, contact := range contacts {
			fmt.Printf("Name: %s, Phone: %s\n", contact.Name, contact.Phone)
		}
	} else if os.Args[1] == "search" && len(os.Args) == 3 {
		contacts := contactbook.SearchContacts(os.Args[2])
		for _, contact := range contacts {
			fmt.Printf("Name: %s, Phone: %s\n", contact.Name, contact.Phone)
		}
	} else {
		PrintUsage()
	}

}

func PrintUsage(){
fmt.Println(`Invalid Command.
		Usage:
		add <name> <phone>   - Add a new contact.
		list                 - List all contacts.
		search <name>        - Search contacts by name.`)
	}
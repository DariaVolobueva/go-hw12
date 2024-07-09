package main

import (
	"fmt"
	"log"
	"os"
	"password-manager/internal/manager"
)

func main() {
	m, err := manager.New("passwords.json")
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: password-manager <command> [arguments]")
		fmt.Println("Commands: list, get <name>, set <name>")
		return
	}

	switch os.Args[1] {
	case "list":
		names := m.List()
		for _, name := range names {
			fmt.Println(name)
		}
	case "get":
		if len(os.Args) != 3 {
			fmt.Println("Usage: password-manager get <name>")
			return
		}
		password, err := m.Get(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(password)
	case "set":
		if len(os.Args) != 3 {
			fmt.Println("Usage: password-manager set <name>")
			return
		}
		fmt.Print("Enter password: ")
		var password string
		fmt.Scanln(&password)
		err := m.Set(os.Args[2], password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Password saved successfully")
	default:
		fmt.Println("Unknown command")
	}
}
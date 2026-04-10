package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// =======================
// DOMAIN
// =======================

type Contact struct {
	ID     int
	Name   string
	Number string
}

// =======================
// REPOSITORY (IN MEMORY)
// =======================

type PhoneBook struct {
	contacts []Contact
	nextID   int
}

func NewPhoneBook(capacity int) *PhoneBook {
	return &PhoneBook{
		contacts: make([]Contact, 0, capacity),
		nextID:   1,
	}
}

// Create
func (p *PhoneBook) Add(name, number string) {
	if p.Exists(name) {
		fmt.Println("⚠️ Contact déjà existant")
		return
	}

	c := Contact{
		ID:     p.nextID,
		Name:   name,
		Number: number,
	}

	p.nextID++
	p.contacts = append(p.contacts, c)
}

// Read all
func (p *PhoneBook) List() []Contact {
	return p.contacts
}

// Search
func (p *PhoneBook) Find(name string) *Contact {
	for i := range p.contacts {
		if strings.EqualFold(p.contacts[i].Name, name) {
			return &p.contacts[i]
		}
	}
	return nil
}

// Delete
func (p *PhoneBook) Delete(name string) bool {
	for i, c := range p.contacts {
		if strings.EqualFold(c.Name, name) {
			p.contacts = append(p.contacts[:i], p.contacts[i+1:]...)
			return true
		}
	}
	return false
}

// Exists
func (p *PhoneBook) Exists(name string) bool {
	return p.Find(name) != nil
}

// =======================
// INPUT LAYER
// =======================

func readInput(reader *bufio.Reader, label string) (string, error) {
	fmt.Print(label)

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)

	if input == "" {
		return "", errors.New("champ vide interdit")
	}

	return input, nil
}

func isExit(input string) bool {
	return strings.EqualFold(input, "finir")
}

// =======================
// BUSINESS
// =======================

func isValidNumber(n string) bool {
	return len(n) >= 6
}

// =======================
// UI
// =======================

func printMenu() {
	fmt.Println("\n===== ANNUAIRE =====")
	fmt.Println("1. Ajouter contact")
	fmt.Println("2. Lister contacts")
	fmt.Println("3. Rechercher contact")
	fmt.Println("4. Supprimer contact")
	fmt.Println("5. Quitter")
	fmt.Print("Choix: ")
}

func printContacts(list []Contact) {
	fmt.Println("\n=== LISTE CONTACTS ===")

	if len(list) == 0 {
		fmt.Println("Aucun contact")
		return
	}

	for _, c := range list {
		fmt.Printf("ID:%d | %s - %s\n", c.ID, c.Name, c.Number)
	}
}

// =======================
// MAIN FLOW
// =======================

func main() {
	reader := bufio.NewReader(os.Stdin)
	phonebook := NewPhoneBook(10)

	for {
		printMenu()

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {

		case "1":
			name, err := readInput(reader, "Nom: ")
			if err != nil {
				fmt.Println("Erreur:", err)
				continue
			}

			if isExit(name) {
				continue
			}

			number, err := readInput(reader, "Numéro: ")
			if err != nil {
				fmt.Println("Erreur:", err)
				continue
			}

			if !isValidNumber(number) {
				fmt.Println("Numéro invalide")
				continue
			}

			phonebook.Add(name, number)

		case "2":
			printContacts(phonebook.List())

		case "3":
			name, _ := readInput(reader, "Nom à rechercher: ")
			c := phonebook.Find(name)

			if c == nil {
				fmt.Println("Contact introuvable")
			} else {
				fmt.Printf("Trouvé: %s - %s\n", c.Name, c.Number)
			}

		case "4":
			name, _ := readInput(reader, "Nom à supprimer: ")
			ok := phonebook.Delete(name)

			if ok {
				fmt.Println("Contact supprimé")
			} else {
				fmt.Println("Contact introuvable")
			}

		case "5":
			fmt.Println("Bye 👋")
			return

		default:
			fmt.Println("Choix invalide")
		}
	}
}

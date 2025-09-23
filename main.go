package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
type contact struct {
	email string
	name  string
}

func (c contact) Display(index int) {
	fmt.Printf("info du contact index : %v | adresse mail %v | nom : %v \n", index, c.email, c.email)
}

func newContact(email string, name string) (contact, error) {
	if email == "" || name == "" {
		return contact{}, errors.New("email et nom ne doivent pas être vides")
	}
	return contact{email: email, name: name}, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Bienvenue sur le mini-crm !\n")

	c := make(map[int]contact)

	add := flag.Bool("add", false, "Ajouter un contact")
	email := flag.String("email", "", "Adresse email du contact")
	name := flag.String("name", "", "Nom du contact")

	flag.Parse()

	if *add {
		if *email == "" || *name == "" {
			fmt.Print("Il manque un champ nécessaire")
		} else {
			index := len(c) + 1
			nContact, _ := newContact(*email, *name)
			c[len(c)+1] = nContact
			c[index].Display(index)
		}

	} else {
		r := true

		for r {

			fmt.Println("Veuillez choisir une option :")
			fmt.Println("1 : Ajouter un contact")
			fmt.Println("2 : Lister tout les contacts")
			fmt.Println("3 : Supprimer un contact")
			fmt.Println("4 : Modifier un contact")
			fmt.Println("5 : Quitter le crm")
			fmt.Println("--------------------------------")

			input, err := reader.ReadString('\n')
			i, _ := strconv.Atoi(strings.TrimSpace(input))

			if err == nil {
				switch i {
				case 1:
					addContact(c, reader)
				case 2:
					displayContact(c)
				case 3:
					deleteContact(c, reader)
				case 4:
					modifyContact(c, reader)
				case 5:
					r = false
				}
			}
		}
	}

}

func addContact(c map[int]contact, reader *bufio.Reader) {
	fmt.Println("Veuillez spécifier une adresse mail:")
	a, err1 := reader.ReadString('\n')
	fmt.Println("Veuillez spécifier un nom:")
	b, err2 := reader.ReadString('\n')
	if err1 == nil && err2 == nil {
		nContact, _ := newContact(a, b)
		c[len(c)+1] = nContact
	}
	fmt.Printf("info du contact ajouté index : %v | adresse mail %v | nom : %v \n", len(c), a, b)
}

func displayContact(c map[int]contact) {
	if len(c) == 0 {
		fmt.Println("Aucun contact à afficher")
	} else {
		for index, val := range c {
			fmt.Printf("index : %v | adresse mail %v | nom : %v \n", index, val.email, val.name)
		}
	}
}

func deleteContact(c map[int]contact, reader *bufio.Reader) {
	fmt.Println("Veuillez spécifier un index à supprimer:")
	input, err1 := reader.ReadString('\n')
	index, _ := strconv.Atoi(strings.TrimSpace(input))
	if err1 == nil {
		_, ok := c[index]
		if ok {
			delete(c, index)
			fmt.Println("Contact supprimé")
		} else {
			fmt.Println("L'index séléctionné ne correspond a aucun contact")
		}
	}
}

func modifyContact(c map[int]contact, reader *bufio.Reader) {
	fmt.Println("Veuillez spécifier un index à modifier:")
	input, err1 := reader.ReadString('\n')
	index, _ := strconv.Atoi(strings.TrimSpace(input))
	if err1 == nil {
		val, ok := c[index]
		if ok {
			fmt.Printf("info du contact à modifier : %v | adresse mail %v | nom : %v \n", len(c), val.email, val.name)
			fmt.Println("Veuillez spécifier une nouvelle adresse mail:")
			a, err2 := reader.ReadString('\n')
			fmt.Println("Veuillez spécifier un nouveau nom:")
			b, err3 := reader.ReadString('\n')
			if err2 == nil && err3 == nil {
				c[index] = contact{
					email: strings.TrimSpace(a),
					name:  strings.TrimSpace(b),
				}
			}
		} else {
			fmt.Println("L'index séléctionné ne correspond a aucun contact")
		}
	}
}

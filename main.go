package main

import (
	"flag"
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	fmt.Printf("Bienvenue sur le mini-crm !\n")

	type contact struct {
		email string
		name  string
	}
	c := make(map[int]contact)

	add := flag.Bool("add", false, "Ajouter un contact")
	email := flag.String("email", "", "Adresse email du contact")
	name := flag.String("name", "", "Nom du contact")

	flag.Parse()

	if *add {
		c[len(c)+1] = contact{
			email: *email,
			name:  *name,
		}

		fmt.Printf("info du contact ajouté index : %v | adresse mail %v | nom : %v \n", len(c), *email, *name)
	} else {
		r := true

		for r {
			var i int
			fmt.Println("Veuillez choisir une option :")
			fmt.Println("1 : Ajouter un contact")
			fmt.Println("2 : Lister tout les contacts")
			fmt.Println("3 : Supprimer un contact")
			fmt.Println("4 : Modifier un contact")
			fmt.Println("5 : Quitter le crm")
			fmt.Println("--------------------------------")
			fmt.Println("")

			_, err := fmt.Scanln(&i)

			if err == nil {
				switch i {
				case 1:
					fmt.Println("Veuillez spécifier une adresse mail:")
					var a string
					_, err1 := fmt.Scanln(&a)
					fmt.Println("Veuillez spécifier un nom:")
					var b string
					_, err2 := fmt.Scanln(&b)
					if err1 == nil && err2 == nil {
						c[len(c)+1] = contact{
							email: a,
							name:  b,
						}
					}
					fmt.Printf("info du contact ajouté index : %v | adresse mail %v | nom : %v \n", len(c), a, b)
				case 2:
					if len(c) == 0 {
						fmt.Print("Aucun contact à afficher")
					} else {
						for index, val := range c {
							fmt.Printf("index : %v | adresse mail %v | nom : %v \n", index, val.email, val.name)
						}
					}
				case 3:
					fmt.Println("Veuillez spécifier un index à supprimer:")
					var index int
					_, err1 := fmt.Scanln(&index)
					if err1 == nil {
						_, ok := c[index]
						if ok {
							delete(c, index)
							fmt.Println("Contact supprimé")
						} else {
							fmt.Println("L'index séléctionné ne correspond a aucun contact")
						}
					}
				case 4:
					fmt.Println("Veuillez spécifier un index à modifier:")
					var index int
					_, err1 := fmt.Scanln(&index)
					if err1 == nil {
						val, ok := c[index]
						if ok {
							fmt.Printf("info du contact à modifier : %v | adresse mail %v | nom : %v \n", len(c), val.email, val.name)
							fmt.Println("Veuillez spécifier une nouvelle adresse mail:")
							var a string
							_, err2 := fmt.Scanln(&a)
							fmt.Println("Veuillez spécifier un nouveau nom:")
							var b string
							_, err3 := fmt.Scanln(&b)
							if err2 == nil && err3 == nil {
								c[index] = contact{
									email: a,
									name:  b,
								}
							}
						} else {
							fmt.Println("L'index séléctionné ne correspond a aucun contact")
						}
					}
				case 5:
					r = false
				}
			}
		}
	}

}

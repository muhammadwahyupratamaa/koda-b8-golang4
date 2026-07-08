package main

import (
	"fmt"
	"koda-b8-Golang4/search"
)

func main() {
	var keyword string
	users := []string{
		"Pra",
		"fajar",
		"Muftin",
		"Bildan",
		"Dimas",
		"Jek",
		"Raply",
		"makruf",
	}

	fmt.Print("Enter the name you want to search for :")
	fmt.Scan(&keyword)
	hasil := search.Search(users,keyword)
	
	if len(hasil)== 0 {
		fmt.Println("User not found")
		return
	}
	fmt.Println("\n Search result :")
	for i, user := range hasil{
		fmt.Printf("%d. %s\n", i+1, user)
	}
}

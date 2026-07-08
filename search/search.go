package search

import (
	"strings"
)

// import "fmt"


func Search(users[]string , keyword string ) []string {
	hasil := []string{

	}
	for _, user := range users{
		if strings.Contains(user,keyword) {
			hasil = append(hasil,user)
		}
	}
		return hasil
}
package main

import (
	"fmt"
	"strings"
)

func main() {
	var link string
	fmt.Println("Masukkan link yang ingin dianalisis:")
	fmt.Scanln(&link)
	fmt.Println("analyze:", link)
	var data = [...]string{"dana.id", "test.aja"}

	// var status bool
	status := false
	for _, check := range data {
		if strings.Contains(link, check) {
			fmt.Println("URL aman")
			status = true
		}
	}
	if !status {
		fmt.Println("Url Berbhaya")
	}
}

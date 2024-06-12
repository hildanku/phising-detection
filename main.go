package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func readWhitelistCSV(filePath string) ([]WhiteList, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Skip header row if present
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	var whitelist []WhiteList
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		whitelist = append(whitelist, WhiteList{Link: record[0]})
	}

	return whitelist, nil
}

type WhiteList struct {
	Link string
}

func main() {

	whitelist, err := readWhitelistCSV("data.csv")
	if err != nil {
		panic(err)
	}

	var link string
	fmt.Println("Masukkan link yang ingin dianalisis:")
	fmt.Scanln(&link)

	isSafe := false
	for _, whitelistItem := range whitelist {
		if strings.Contains(link, whitelistItem.Link) {
			isSafe = true
			break
		}
	}

	if isSafe {
		fmt.Println("URL aman")
	} else {
		fmt.Println("URL Berbahaya")
	}
}

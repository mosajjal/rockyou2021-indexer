package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func checkString(input string, indexDir string) bool {
	lineInput := input
	lineBytes := []byte(lineInput)
	listPath := fmt.Sprintf("%v/%v/%v/%v", indexDir, lineBytes[0], lineBytes[1], lineBytes[2])
	f, err := os.OpenFile(fmt.Sprintf("%v/list", listPath), os.O_RDONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	fScanner := bufio.NewScanner(f)
	for fScanner.Scan() {
		if lineInput == fScanner.Text() {
			return true
		}
	}
	if err := fScanner.Err(); err != nil {
		log.Fatal(err)
	}
	return false
}

func main() {
	curDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	var indexDir = curDir + "/indexed"

	if len(os.Args) < 2 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lineInput := scanner.Text()
			fmt.Printf("%v: %v\n", lineInput, checkString(lineInput, indexDir))
		}
	} else {
		lineInput := os.Args[1]
		fmt.Printf("%v: %v\n", lineInput, checkString(lineInput, indexDir))
	}
}

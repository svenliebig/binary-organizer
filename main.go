package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter the path to the file: ")

	// arguments
	if len(os.Args) > 1 {
		fmt.Println("Arguments: ", os.Args[1:])
	}

	var path string
	_, err := fmt.Scanln(&path)
	if err != nil {
		fmt.Println("Error reading the path")
		return
	}

	err = os.WriteFile(".path", []byte(fmt.Sprintf("PATH='%s'", path)), 0644)

	if err != nil {
		fmt.Println("Error writing the file")
		return
	}
}

type EnvironmentVariable struct {
	Name  string
	Value string
}

package main

import (
	"fmt"
	"os"

	"github.com/svenliebig/binary-organizer/internal/path"
)

func main() {
	fmt.Println("Enter the path to the file: ")

	// arguments
	if len(os.Args) > 1 {
		fmt.Println("Arguments: ", os.Args[1:])
	}

	// cwd
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting the current working directory")
		return
	}

	fmt.Println("Current working directory: ", cwd)

	// pwd
	pwd, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting the user home directory")
		return
	}

	fmt.Println("User home directory: ", pwd)

	p := path.NewPathVariable()

	fmt.Println(p.Export())

	// err = os.WriteFile(".path", []byte(fmt.Sprintf("PATH='%s'", path)), 0644)
	//
	// if err != nil {
	// 	fmt.Println("Error writing the file")
	// 	return
	// }
}

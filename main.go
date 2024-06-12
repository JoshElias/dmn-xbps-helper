package main

import (
	// "bufio"
	"fmt"
	// "log"
	"os"
	// "path/filepath"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting cwd:", err)
		return
	}
	fmt.Println("CWD:", cwd)

	// f, err := os.Open('

}

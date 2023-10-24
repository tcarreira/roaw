package main

import (
	"fmt"
	"os"
)

var (
	version = "dev"
	commit  = ""
	dateStr = ""
)

func main() {
	fmt.Println("Hello from roaw")

	if len(os.Args) > 1 && os.Args[1] == "-v" {
		fmt.Printf("%s (%s - %s)\n", version, commit, dateStr)
	}
}

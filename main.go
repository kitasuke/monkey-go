package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/kitasuke/monkey-go/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		print(err)
	}

	fmt.Printf("Hello %s! This is the Monky programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
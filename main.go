package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/adamerikoff/ponGo/src/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the ponGo programming language!\n",
		user.Username)
	fmt.Printf("Type your commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

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
	fmt.Printf("Hallo %s! Das ist ponGo!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}

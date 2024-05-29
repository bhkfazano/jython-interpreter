package main

import (
	"fmt"
	"kogab-interpreter/pkg/repl"
	"os"
	"os/user"
)

func main() {
	var user, err = user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("[%s] Welcome to the Jython console\n", user.Username)

	repl.Start(os.Stdin, os.Stdout, user.Username)
}

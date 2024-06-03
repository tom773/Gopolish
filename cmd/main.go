package main

import (
	"fmt"
	"os"
	pretty "tmpretty/utils"
)

func main() {
	err := pretty.Pretty(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

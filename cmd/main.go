package main

import (
	"fmt"
	"os"

	pretty "github.com/tom773/gopolish/utils"
)

func main() {
	err := pretty.Pretty(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

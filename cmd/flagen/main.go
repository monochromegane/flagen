package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/monochromegane/flagen"
)

func main() {
	err := flagen.Run(os.Args[1:], os.Stdout, os.Stderr)
	if err != nil && err != flag.ErrHelp {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

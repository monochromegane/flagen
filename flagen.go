// Package flagen provides a command line option parser generator using command line option.
package flagen

import (
	"errors"
	"flag"
	"io"
)

// Run runs the flagen.
func Run(args []string, outStream, errStream io.Writer) error {
	fs := flag.NewFlagSet("flagen", flag.ContinueOnError)
	fs.SetOutput(errStream)
	if err := fs.Parse(args); err != nil {
		return err
	}

	args = fs.Args()
	if len(args) < 1 {
		return errors.New("template is required")
	}

	tmpl, err := NewTemplate(args[0])
	if err != nil {
		return err
	}
	return tmpl.Execute(outStream, args[1:])
}

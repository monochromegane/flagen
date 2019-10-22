package flagen

import (
	"errors"
	"flag"
	"io"
)

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

	gen := &generator{template: args[0]}
	return gen.run(args[1:], outStream, errStream)
}

package flagen

import (
	"io"
)

type generator struct {
	outStream io.Writer

	template string
}

func (g *generator) run(args []string, outStream, errStream io.Writer) error {
	g.outStream = outStream

	flagSet := &FlagSet{}
	err := flagSet.Parse(args)
	if err != nil {
		return err
	}
	if err := g.generate(flagSet.Flags()); err != nil {
		return err
	}
	return nil
}

func (g *generator) generate(flags []Flag) error {
	return nil
}

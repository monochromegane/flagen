package flagen

import (
	"io"
)

type generator struct {
	outStream io.Writer

	template string
	args     []string
	flags    []flagSet
}

func (g *generator) run(args []string, outStream, errStream io.Writer) error {
	g.outStream = outStream
	g.args = args

	if err := g.parse(); err != nil {
		return err
	}
	if err := g.generate(); err != nil {
		return err
	}
	return nil
}

func (g *generator) parse() error {
	for {
		seen, err := g.parseOne()
		if seen {
			continue
		}
		return err
	}
}

func (g *generator) parseOne() (bool, error) {
	return false, nil
}

func (g *generator) generate() error {
	return nil
}

package flagen

import (
	"io"
)

type generator struct {
	template string
}

func (g *generator) run(args []string, outStream, errStream io.Writer) error {
	fs := &FlagSet{}
	if err := fs.Parse(args); err != nil {
		return err
	}

	tmpl, err := templateFrom(g.template)
	if err != nil {
		return err
	}
	return tmpl.Execute(outStream, map[string]interface{}{
		"Flags": fs.Flags(),
		"Args":  fs.Args(),
	})
}

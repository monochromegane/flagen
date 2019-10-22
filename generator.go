package flagen

import (
	"fmt"
	"io"
	"text/template"
)

type generator struct {
	outStream io.Writer

	template string
}

func (g *generator) run(args []string, outStream, errStream io.Writer) error {
	g.outStream = outStream

	fs := &FlagSet{}
	if err := fs.Parse(args); err != nil {
		return err
	}
	if err := g.generate(fs.Flags()); err != nil {
		return err
	}
	return nil
}

func (g *generator) generate(flags []Flag) error {
	fn, ok := templateMap[g.template]
	if !ok {
		return fmt.Errorf("template dosen't exist: %s", g.template)
	}

	tmpl := template.New("flagen")
	tmpl = tmpl.Funcs(templateFuncMap)
	tmpl, err := tmpl.Parse(fn())
	if err != nil {
		return err
	}
	return tmpl.Execute(g.outStream, flags)
}

package flagen

import (
	"fmt"
	"io"
	"text/template"
)

// NewTemplate creates a new Template.
func NewTemplate(path string) (*Template, error) {
	tmpl, err := templateFrom(path)
	if err != nil {
		return nil, err
	}
	return &Template{tmpl: tmpl}, nil
}

// Template is the representation of a parsed template.
type Template struct {
	tmpl *template.Template
}

// Execute applies a parsed template to the specified arguments, and writes the output to outStream.
func (t *Template) Execute(outStream io.Writer, args []string) error {
	fs := &flagSet{}
	if err := fs.parse(args); err != nil {
		return err
	}

	return t.tmpl.Execute(outStream, map[string]interface{}{
		"Flags": fs.flags(),
		"Args":  fs.args(),
	})
}

func templateFrom(loc string) (*template.Template, error) {
	if tmpl, err := templateFromFile(loc); err == nil {
		return tmpl, nil
	} else if tmpl, err := templateFromPreset(loc); err != nil {
		return nil, err
	} else {
		return tmpl, nil
	}
}

func templateFromFile(path string) (*template.Template, error) {
	tmpl := template.New(path)
	tmpl = tmpl.Funcs(TemplateFuncMap)
	return tmpl.ParseFiles(path)
}

func templateFromPreset(key string) (*template.Template, error) {
	fn, ok := templateMap[key]
	if !ok {
		return nil, fmt.Errorf("template dosen't exist: %s", key)
	}
	tmpl := template.New(key)
	tmpl = tmpl.Funcs(TemplateFuncMap)
	return tmpl.Parse(fn())
}

package flagen

import (
	"fmt"
	"text/template"

	"github.com/iancoleman/strcase"
)

var TemplateFuncMap template.FuncMap = template.FuncMap{
	"ToSnake":              strcase.ToSnake,
	"ToScreamingSnake":     strcase.ToScreamingSnake,
	"ToKebab":              strcase.ToKebab,
	"ToScreamingKebab":     strcase.ToScreamingKebab,
	"ToDelimited":          strcase.ToDelimited,
	"ToScreamingDelimited": strcase.ToScreamingDelimited,
	"ToCamel":              strcase.ToCamel,
	"ToLowerCamel":         strcase.ToLowerCamel,
}

type templateTextFunc func() string

var templateMap map[string]templateTextFunc = map[string]templateTextFunc{
	"go": templateTextGo,
}

func templateTextGo() string {
	return `var (
{{ range $flag := .Flags }}	{{ ToLowerCamel $flag.Name }}	{{ if eq "float" $flag.Value.Type -}}
	float64
{{ else -}}
	{{$flag.Value.Type}}
{{ end -}}
{{ end }})

func init() {
{{ range $flag := .Flags }}	{{ $type := $flag.Value.Type -}}
{{ if eq "float" $type -}}{{ $type = "float64" -}}{{ end -}}
{{ if eq "string" $type -}}
	flag.{{ ToCamel $type }}Var(&{{ ToLowerCamel $flag.Name }}, "{{ $flag.Name }}", "{{ $flag.Value.Get }}", "usage of {{ $flag.Name }}")
{{ else -}}
	flag.{{ ToCamel $type }}Var(&{{ ToLowerCamel $flag.Name }}, "{{ $flag.Name }}", {{ $flag.Value.Get }}, "usage of {{ $flag.Name }}")
{{ end -}}
{{ end }}}
`
}

func templateFrom(loc string) (*template.Template, error) {
	if tmpl, err := TemplateFromFile(loc); err == nil {
		return tmpl, nil
	} else if tmpl, err := templateFromPreset(loc); err != nil {
		return nil, err
	} else {
		return tmpl, nil
	}
}

func TemplateFromFile(path string) (*template.Template, error) {
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

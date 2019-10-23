package flagen

import (
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
	"py": templateTextPython,
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

func templateTextPython() string {
	return `import argparse
parser = argparse.ArgumentParser()
{{ range $flag := .Flags }}{{ $type := $flag.Value.Type -}}

{{ $prefix := "--" -}}{{ if eq (len $flag.Name) 1 -}}{{ $prefix = "-" -}}{{ end -}}

{{ $type := "" -}}
{{ if or ( eq "int" $flag.Value.Type ) ( eq "float" $flag.Value.Type ) -}}{{ $type = $flag.Value.Type }}{{ end -}}

{{ $default := "" -}}
{{ if eq "string" $flag.Value.Type -}}{{ $default = printf "\"%s\"" $flag.Value.Get }}{{ end -}}
{{ if or ( eq "int" $flag.Value.Type ) ( eq "float" $flag.Value.Type ) -}}{{ $default = $flag.Value.Get }}{{ end -}}

{{ $action := "" -}}
{{ if eq "bool" $flag.Value.Type -}}{{ $action = printf "\"store_%t\"" $flag.Value.Get }}{{ end -}}

parser.add_argument("{{ $prefix }}{{ $flag.Name }}"{{ with $type }}, type={{.}}{{ end -}}{{ with $default }}, default={{.}}{{ end -}}{{ with $action }}, action={{ $action }}{{ end -}}, help="Help of {{ $flag.Name }}")
{{ end -}}

{{ range $arg := .Args -}}
parser.add_argument("{{ $arg }}", help="Help of {{ $arg }}")
{{ end -}}
`
}

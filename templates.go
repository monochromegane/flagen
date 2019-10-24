package flagen

import (
	"text/template"

	"github.com/iancoleman/strcase"
)

// TemplateFuncMap has template functions for converting string case.
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
	"rb": templateTextRuby,
	"sh": templateTextShell,
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
args = parser.parse_args()
`
}

func templateTextRuby() string {
	return `require 'optparse'

opts = {
{{ range $flag := .Flags }}  {{ $default := $flag.Value.Get -}}
{{ if eq "string" $flag.Value.Type -}}{{ $default = printf "'%s'" $flag.Value.Get }}{{ end -}}
  {{ ToSnake $flag.Name -}}: {{ $default -}},
{{ end -}}
}

OptionParser.new do |op|
{{ range $flag := .Flags }}  {{ $type := $flag.Value.Type -}}
{{ $prefix := "--" -}}{{ if eq (len $flag.Name) 1 -}}{{ $prefix = "-" -}}{{ end -}}
{{ $postfix := "" -}}{{ if ne "bool" $flag.Value.Type -}}{{ $postfix = "VALUE" -}}{{ end -}}
{{ $cast := "" -}}
{{ if eq "int"   $flag.Value.Type -}}{{ $cast = "to_i" }}{{ end -}}
{{ if eq "float" $flag.Value.Type -}}{{ $cast = "to_f" }}{{ end -}}
  op.on('{{ $prefix }}{{ $flag.Name }}{{ with $postfix }} [{{.}}]{{ end }}', 'Desc of {{ $flag.Name }}') {|v| opts[:{{ ToSnake $flag.Name }}] = v{{ with $cast }}.{{.}}{{ end }} }
{{ end }}
  op.parse!(ARGV)
end
`
}

func templateTextShell() string {
	return `{{ range $flag := .Flags -}}
{{ if eq ( len $flag.Name ) 1 -}}
{{ $value := $flag.Value.Get -}}
{{ if eq "bool" $flag.Value.Type -}}{{ $value = ToScreamingSnake ( printf "%t" $flag.Value.Get ) }}{{ end -}}
{{ ToScreamingSnake $flag.Name }}="{{ $value }}"
{{ end -}}
{{ end }}
while getopts {{ range $flag := .Flags -}}
{{ if eq ( len $flag.Name ) 1 -}}
{{ $postfix := "" -}}
{{ if ne "bool" $flag.Value.Type -}}{{ $postfix = ":" }}{{ end -}}
{{ $flag.Name }}{{ with $postfix}}{{.}}{{ end -}}
{{ end }}
{{- end }} OPT
do
  case $OPT in{{ range $flag := .Flags }}{{ if eq ( len $flag.Name ) 1 }}
    {{ $value := "$OPTARG" -}}
    {{ if eq "bool" $flag.Value.Type -}}{{ $value = ToScreamingSnake ( printf "%t" ( eq false $flag.Value.Get ) ) -}}{{ end -}}
    "{{ $flag.Name }}" ) {{ ToScreamingSnake $flag.Name }}="{{ $value }}";;
    {{- end -}}
    {{ end }}
  esac
done

` + "shift `expr $OPTIND - 1`\n"
}

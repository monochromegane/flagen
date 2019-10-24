# flagen

A command line option parser generator using command line option.
It generates a useful boilerplate of option parser for various programing languages using the format of command line options that will be used.

## Usage

```sh
$ flagen YOUR_TEMPLATE YOUR_COMMAND_LINE_OPTIONS...
```

```sh
$ flagen go -i 1 -f 1.1 -s abc -b1 -b2=true arg1 arg2
var (
        i       int
        f       float64
        s       string
        b1      bool
        b2      bool
)

func init() {
        flag.IntVar(&i, "i", 1, "usage of i")
        flag.Float64Var(&f, "f", 1.1, "usage of f")
        flag.StringVar(&s, "s", "abc", "usage of s")
        flag.BoolVar(&b1, "b1", false, "usage of b1")
        flag.BoolVar(&b2, "b2", true, "usage of b2")
}
```

## Templates

Flagen has preset templates for {go,py,rb,sh}.
You can specify them as template name.

Of course you can also specify your template file path.

The template is parsed as [text/template](https://golang.org/pkg/text/template/) of Go.

In your template, you can use `.Flags` and `.Args`.
`Flags` has some `Flag` that has `Name` and `Value` of each option.
And `Value` has `Get` and `Type` method.
`Get` returns the value of option.
`Type` returns the estimated type of option (`int`, `float`, `string`, `bool`).
`Args` has some string which is argument.

And you can use functions for converting string case.
It provided from [iancoleman/strcase](https://github.com/iancoleman/strcase).

## Collaboration

### Vim

In your source which is being opened by Vim, you can insert the boilerplate.

```
:r!flagen YOUR_TEMPLATE YOUR_COMMAND_LINE_OPTIONS...
```

### Your boilerplate tool

Flagen provides generator as library.
You can use it in your boilerplate tool in Go as the following,

```go
	tmpl, err := flagen.NewTemplate(args[0])
	if err != nil {
		return err
	}
	return tmpl.Execute(outStream, args[1:])
```

## Workarounds

### Ambiguous flag

Flagen consider the flag which has no value as bool type.
If you specify bool flag and argument as the following,

```sh
$ flagen TEMPLATE --bool-flag arg1
```

it consider the flag that has string value.

When you want to avoid the case, you can use `=` as the following,

```sh
$ flagen TEMPLATE --bool-flag=false arg1
```

## Installation

```sh
$ go get github.com/monochromegane/flagen/...
```

## License

[MIT](https://github.com/monochromegane/flagen/blob/master/LICENSE)

## Author

[monochromegane](https://github.com/monochromegane)

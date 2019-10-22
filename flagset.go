package flagen

import (
	"container/list"
	"fmt"
)

type FlagSet struct {
	args    []string
	flagSet *flagSet
}

func (fs *FlagSet) Parse(args []string) error {
	fs.args = args
	return fs.parse()
}

func (fs *FlagSet) Arg(i int) string {
	if len(fs.args) <= i {
		return ""
	}
	return fs.args[i]
}

func (fs *FlagSet) Args() []string {
	return fs.args
}

func (fs *FlagSet) Flag(i int) Flag {
	flags := fs.Flags()
	if len(flags) <= i {
		return Flag{}
	}
	return flags[i]
}

func (fs *FlagSet) Flags() []Flag {
	return fs.flagSet.toList()
}

func (fs *FlagSet) parse() error {
	fs.flagSet = newFlagSet()
	for {
		seen, err := fs.parseOne()
		if seen {
			continue
		}
		return err
	}
}

func (fs *FlagSet) parseOne() (bool, error) {
	if len(fs.args) == 0 {
		return false, nil
	}
	s := fs.args[0]
	if len(s) < 2 || s[0] != '-' {
		return false, nil
	}

	numMinuses := 1
	if s[1] == '-' {
		numMinuses++
		if len(s) == 2 {
			fs.args = fs.args[1:]
			return false, nil
		}
	}
	name := s[numMinuses:]
	if len(name) == 0 || name[0] == '-' || name[0] == '=' {
		return false, fmt.Errorf("bad flag syntax: %s", s)
	}

	fs.args = fs.args[1:]
	hasValue := false
	value := ""
	for i := 1; i < len(name); i++ {
		if name[i] == '=' {
			value = name[i+1:]
			hasValue = true
			name = name[0:i]
			break
		}
	}

	if !hasValue && len(fs.args) > 0 {
		if fs.args[0][0] != '-' {
			hasValue = true
			value, fs.args = fs.args[0], fs.args[1:]
		}
	}

	if !hasValue {
		value = "false"
	}

	fs.flagSet.set(name, toValue(value))

	return true, nil
}

func toValue(val string) value {
	if v, err := newIntValue(val); err == nil {
		return v
	}
	if v, err := newFloatValue(val); err == nil {
		return v
	}
	if v, err := newBoolValue(val); err == nil {
		return v
	}
	v, _ := newStringValue(val)
	return v
}

func newFlagSet() *flagSet {
	return &flagSet{flags: list.New()}
}

type flagSet struct {
	flags *list.List
}

func (fs *flagSet) set(name string, value value) {
	if e := fs.lookup(name); e != nil {
		fs.flags.Remove(e)
	}
	fs.flags.PushBack(Flag{Name: name, Value: value})
}

func (fs *flagSet) lookup(name string) *list.Element {
	for e := fs.flags.Front(); e != nil; e = e.Next() {
		if f, ok := e.Value.(Flag); ok && f.Name == name {
			return e
		}
	}
	return nil
}

func (fs *flagSet) toList() []Flag {
	flags := make([]Flag, fs.flags.Len())
	var i int
	for e := fs.flags.Front(); e != nil; e = e.Next() {
		flags[i] = e.Value.(Flag)
		i++
	}
	return flags
}

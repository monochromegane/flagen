package flagen

import (
	"container/list"
	"fmt"
)

type flagSet struct {
	argList []string
	flagSet *orderedFlagSet
}

func (fs *flagSet) parse(args []string) error {
	fs.argList = args
	fs.flagSet = newOrderedFlagSet()
	for {
		seen, err := fs.parseOne()
		if seen {
			continue
		}
		return err
	}
}

func (fs *flagSet) arg(i int) string {
	if len(fs.argList) <= i {
		return ""
	}
	return fs.argList[i]
}

func (fs *flagSet) args() []string {
	return fs.argList
}

func (fs *flagSet) flag(i int) Flag {
	flags := fs.flags()
	if len(flags) <= i {
		return Flag{}
	}
	return flags[i]
}

func (fs *flagSet) flags() []Flag {
	return fs.flagSet.toList()
}

func (fs *flagSet) parseOne() (bool, error) {
	if len(fs.argList) == 0 {
		return false, nil
	}
	s := fs.argList[0]
	if len(s) < 2 || s[0] != '-' {
		return false, nil
	}

	numMinuses := 1
	if s[1] == '-' {
		numMinuses++
		if len(s) == 2 {
			fs.argList = fs.argList[1:]
			return false, nil
		}
	}
	name := s[numMinuses:]
	if len(name) == 0 || name[0] == '-' || name[0] == '=' {
		return false, fmt.Errorf("bad flag syntax: %s", s)
	}

	fs.argList = fs.argList[1:]
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

	if !hasValue && len(fs.argList) > 0 {
		if fs.argList[0][0] != '-' {
			hasValue = true
			value, fs.argList = fs.argList[0], fs.argList[1:]
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

func newOrderedFlagSet() *orderedFlagSet {
	return &orderedFlagSet{flags: list.New()}
}

type orderedFlagSet struct {
	flags *list.List
}

func (fs *orderedFlagSet) set(name string, value value) {
	if e := fs.lookup(name); e != nil {
		fs.flags.Remove(e)
	}
	fs.flags.PushBack(Flag{Name: name, Value: value})
}

func (fs *orderedFlagSet) lookup(name string) *list.Element {
	for e := fs.flags.Front(); e != nil; e = e.Next() {
		if f, ok := e.Value.(Flag); ok && f.Name == name {
			return e
		}
	}
	return nil
}

func (fs *orderedFlagSet) toList() []Flag {
	flags := make([]Flag, fs.flags.Len())
	var i int
	for e := fs.flags.Front(); e != nil; e = e.Next() {
		flags[i] = e.Value.(Flag)
		i++
	}
	return flags
}

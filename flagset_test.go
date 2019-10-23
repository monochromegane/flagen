package flagen

import (
	"strconv"
	"testing"
)

func TestToValue(t *testing.T) {
	tests := map[string]string{
		"1":    "int",
		"1.0":  "float",
		"true": "bool",
		"abc":  "string",
	}
	for test, expect := range tests {
		value := toValue(test)
		if typ := value.Type(); typ != expect {
			t.Errorf("toValue should return type %s, but %s\n", expect, typ)
		}
	}
}

func TestFlagSet(t *testing.T) {
	flagSet := newOrderedFlagSet()

	idx := []int64{0, 1, 2, 3, 4}
	names := []string{"a", "b", "c", "b", "a"}
	for i, name := range names {
		v, _ := newIntValue(strconv.Itoa(i))
		flagSet.set(name, v)
	}

	flags := flagSet.toList()
	if len(flags) != 3 {
		t.Errorf("flagSet.toList() should return 3 flags, but %d\n", len(flags))
	}
	for i, f := range flags {
		if name := f.Name; name != names[i+2] {
			t.Errorf("flags[%d] should have name %s, but %s\n", i, names[i+2], name)
		}
		if typ := f.Value.Type(); typ != "int" {
			t.Errorf("flags[%d] should have type %s, but %s\n", i, "int", typ)
		}
		if value := f.Value.Get(); value.(int64) != idx[i+2] {
			t.Errorf("flags[%d] should have value %d, but %d\n", i, idx[i+2], value)
		}
	}
}

func TestFlagSetParse(t *testing.T) {
	for _, arguments := range [][]string{
		[]string{"-i", "1", "-f", "1.0", "-s", "abc", "-b1", "-b2=true", "arg1", "arg2"},
		[]string{"--i", "1", "--f", "1.0", "--s", "abc", "--b1", "--b2=true", "arg1", "arg2"},
	} {
		flagSet := &flagSet{}
		err := flagSet.parse(arguments)
		if err != nil {
			t.Errorf("flagSet.Parse should not return err, but %v\n", err)
		}

		flags := flagSet.flags()
		if len(flags) != 5 {
			t.Errorf("flagSet.Parse should not return 5 flags, but %d\n", len(flags))
		}

		if f := flags[0]; f.Name != "i" || f.Value.Type() != "int" || f.Value.Get().(int64) != 1 {
			t.Errorf("flagSet.Parse should not return flag of int\n")
		}
		if f := flags[1]; f.Name != "f" || f.Value.Type() != "float" || f.Value.Get().(float64) != 1.0 {
			t.Errorf("flagSet.Parse should not return flag of float\n")
		}
		if f := flags[2]; f.Name != "s" || f.Value.Type() != "string" || f.Value.Get().(string) != "abc" {
			t.Errorf("flagSet.Parse should not return flag of float\n")
		}
		if f := flags[3]; f.Name != "b1" || f.Value.Type() != "bool" || f.Value.Get().(bool) != false {
			t.Errorf("flagSet.Parse should not return flag of bool\n")
		}
		if f := flags[4]; f.Name != "b2" || f.Value.Type() != "bool" || f.Value.Get().(bool) != true {
			t.Errorf("flagSet.Parse should not return flag of bool\n")
		}

		args := flagSet.args()
		if len(args) != 2 {
			t.Errorf("flagSet.Parse should not return 2 args, but %d\n", len(args))
		}

		if arg := args[0]; arg != "arg1" {
			t.Errorf("flagSet.Parse should not return arg %s, but %s\n", "arg1", arg)
		}
		if arg := args[1]; arg != "arg2" {
			t.Errorf("flagSet.Parse should not return arg %s, but %s\n", "arg2", arg)
		}
	}
}

func TestFlagSetParseBadSyntax(t *testing.T) {
	for _, arguments := range [][]string{
		[]string{"---", "1"},
		[]string{"--=", "1"},
	} {
		flagSet := &flagSet{}
		err := flagSet.parse(arguments)

		if err == nil {
			t.Errorf("flagSet.Parse should return error\n")
		}
	}
}

func TestFlagSetParseTerminate(t *testing.T) {
	for i, arguments := range [][]string{
		[]string{"-i", "1", "-", "-f", "1.0", "arg1"},
		[]string{"-i", "1", "--", "-f", "1.0", "arg1"},
	} {
		flagSet := &flagSet{}
		err := flagSet.parse(arguments)
		if err != nil {
			t.Errorf("flagSet.Parse should not return err, but %v\n", err)
		}

		flags := flagSet.flags()
		if len(flags) != 1 {
			t.Errorf("flagSet.Parse should not return 5 flags, but %d\n", len(flags))
		}

		args := flagSet.args()
		if len(args) != 4-i {
			t.Errorf("flagSet.Parse should return %d args, but %d\n", 4-i, len(args))
		}
		expect := arguments[2+i:]
		for j, _ := range args {
			if args[j] != expect[j] {
				t.Errorf("flagSet.Parse should return args %v, but %v\n", arguments[2+i:], args)
			}
		}
	}
}

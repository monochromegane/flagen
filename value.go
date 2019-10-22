package flagen

import "strconv"

type value interface {
	set(string) error
	Get() interface{}
	Type() string
}

func newBoolValue(v string) (*boolValue, error) {
	bv := &boolValue{}
	err := bv.set(v)
	return bv, err
}

type boolValue struct {
	v bool
}

func (b *boolValue) set(s string) error {
	v, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	b.v = v
	return err
}

func (b *boolValue) Get() interface{} { return b.v }
func (b *boolValue) Type() string     { return "bool" }

func newIntValue(v string) (*intValue, error) {
	iv := &intValue{}
	err := iv.set(v)
	return iv, err
}

type intValue struct {
	v int64
}

func (i *intValue) set(s string) error {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	if err != nil {
		return err
	}
	i.v = v
	return err
}

func (i *intValue) Get() interface{} { return i.v }
func (i *intValue) Type() string     { return "int" }

func newFloatValue(v string) (*floatValue, error) {
	fv := &floatValue{}
	err := fv.set(v)
	return fv, err
}

type floatValue struct {
	v float64
}

func (f *floatValue) set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	f.v = v
	return err
}

func (f *floatValue) Get() interface{} { return f.v }
func (f *floatValue) Type() string     { return "float" }

func newStringValue(v string) (*stringValue, error) {
	sv := &stringValue{}
	err := sv.set(v)
	return sv, err
}

type stringValue struct {
	v string
}

func (s *stringValue) set(val string) error {
	s.v = val
	return nil
}

func (s *stringValue) Get() interface{} { return s.v }
func (s *stringValue) Type() string     { return "string" }

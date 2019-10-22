package flagen

import "testing"

func TestBoolValue(t *testing.T) {
	expect := "true"
	v, err := newBoolValue(expect)
	if err != nil {
		t.Errorf("newBoolValue should not return error, but %v\n", err)
	}
	b, ok := v.Get().(bool)
	if !ok {
		t.Errorf("boolValue.Get should return bool\n")
	}
	if !b {
		t.Errorf("boolValue.Get should return %s, but %t\n", expect, b)
	}
	if typ := v.Type(); typ != "bool" {
		t.Errorf("boolValue.Type should return bool, but %s\n", typ)
	}
}

func TestIntValue(t *testing.T) {
	expect := "10"
	v, err := newIntValue(expect)
	if err != nil {
		t.Errorf("newIntValue should not return error, but %v\n", err)
	}
	i, ok := v.Get().(int64)
	if !ok {
		t.Errorf("intValue.Get should return int\n")
	}
	if i != 10 {
		t.Errorf("intValue.Get should return %s, but %d\n", expect, i)
	}
	if typ := v.Type(); typ != "int" {
		t.Errorf("intValue.Type should return int, but %s\n", typ)
	}
}

func TestFloatValue(t *testing.T) {
	expect := "10.0"
	v, err := newFloatValue(expect)
	if err != nil {
		t.Errorf("newFloatValue should not return error, but %v\n", err)
	}
	f, ok := v.Get().(float64)
	if !ok {
		t.Errorf("floatValue.Get should return float\n")
	}
	if f != 10.0 {
		t.Errorf("floatValue.Get should return %s, but %f\n", expect, f)
	}
	if typ := v.Type(); typ != "float" {
		t.Errorf("floatValue.Type should return float, but %s\n", typ)
	}
}

func TestStringValue(t *testing.T) {
	expect := "abc"
	v, err := newStringValue(expect)
	if err != nil {
		t.Errorf("newStringValue should not return error, but %v\n", err)
	}
	s, ok := v.Get().(string)
	if !ok {
		t.Errorf("stringValue.Get should return string\n")
	}
	if s != expect {
		t.Errorf("stringValue.Get should return %s, but %s\n", expect, s)
	}
	if typ := v.Type(); typ != "string" {
		t.Errorf("stringValue.Type should return string, but %s\n", typ)
	}
}

package flagen

import "os"

func ExampleTemplate_Execute_Go() {
	tmpl, _ := NewTemplate("go")

	args := []string{"-i", "1", "-f", "1.1", "-s", "abc", "-b1", "-b2=true", "arg1", "arg2"}
	tmpl.Execute(os.Stdout, args)

	// Output:
	// var (
	// 	i	int
	// 	f	float64
	// 	s	string
	// 	b1	bool
	// 	b2	bool
	// )
	//
	// func init() {
	// 	flag.IntVar(&i, "i", 1, "usage of i")
	// 	flag.Float64Var(&f, "f", 1.1, "usage of f")
	// 	flag.StringVar(&s, "s", "abc", "usage of s")
	// 	flag.BoolVar(&b1, "b1", false, "usage of b1")
	// 	flag.BoolVar(&b2, "b2", true, "usage of b2")
	// }
}

func ExampleTemplate_Execute_Python() {
	tmpl, _ := NewTemplate("py")

	args := []string{"-i", "1", "-f", "1.1", "-s", "abc", "-b1", "-b2=true", "arg1", "arg2"}
	tmpl.Execute(os.Stdout, args)

	// Output:
	// import argparse
	// parser = argparse.ArgumentParser()
	// parser.add_argument("-i", type=int, default=1, help="Help of i")
	// parser.add_argument("-f", type=float, default=1.1, help="Help of f")
	// parser.add_argument("-s", default="abc", help="Help of s")
	// parser.add_argument("--b1", action="store_false", help="Help of b1")
	// parser.add_argument("--b2", action="store_true", help="Help of b2")
	// parser.add_argument("arg1", help="Help of arg1")
	// parser.add_argument("arg2", help="Help of arg2")
}

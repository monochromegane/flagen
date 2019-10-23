package flagen

import "os"

func ExampleGenerator_Run_Go() {
	gen := &generator{template: "go"}

	args := []string{"-i", "1", "-f", "1.1", "-s", "abc", "-b1", "-b2=true", "arg1", "arg2"}
	gen.run(args, os.Stdout, os.Stderr)

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

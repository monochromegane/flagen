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
	// args = parser.parse_args()
}

func ExampleTemplate_Execute_Ruby() {
	tmpl, _ := NewTemplate("rb")

	args := []string{"-i", "1", "-f", "1.1", "-s", "abc", "-b1", "-b2=true", "arg1", "arg2"}
	tmpl.Execute(os.Stdout, args)

	// Output:
	// require 'optparse'
	//
	// opts = {
	//   i: 1,
	//   f: 1.1,
	//   s: 'abc',
	//   b_1: false,
	//   b_2: true,
	// }
	//
	// OptionParser.new do |op|
	//   op.on('-i [VALUE]', 'Desc of i') {|v| opts[:i] = v.to_i }
	//   op.on('-f [VALUE]', 'Desc of f') {|v| opts[:f] = v.to_f }
	//   op.on('-s [VALUE]', 'Desc of s') {|v| opts[:s] = v }
	//   op.on('--b1', 'Desc of b1') {|v| opts[:b_1] = v }
	//   op.on('--b2', 'Desc of b2') {|v| opts[:b_2] = v }
	//
	//   op.parse!(ARGV)
	// end
}

func ExampleTemplate_Execute_Shell() {
	tmpl, _ := NewTemplate("sh")

	args := []string{"-i", "1", "-s", "abc", "-b", "-c=true", "arg1", "arg2"}
	tmpl.Execute(os.Stdout, args)

	// Output:
	// I="1"
	// S="abc"
	// B="FALSE"
	// C="TRUE"
	//
	// while getopts i:s:bc OPT
	// do
	//   case $OPT in
	//     "i" ) I="$OPTARG";;
	//     "s" ) S="$OPTARG";;
	//     "b" ) B="TRUE";;
	//     "c" ) C="FALSE";;
	//   esac
	// done
	//
	// shift `expr $OPTIND - 1`
}

package flags

import (
	"flag"
)

type Options struct {
	Color     string
	Output    string
	Substring string
	Input     string
	Banner    string
}

func ParseOptions() Options {
	var options Options
	flag.StringVar(&options.Color, "color", "", "Usage: --color=<color>")
	flag.StringVar(&options.Output, "output", "", "Usage: --ouput=<file.txt>")
	flag.Parse()

	switch len(flag.Args()) {
	case 1:
		options.Input = flag.Arg(0)
	case 2:
		options.Input = flag.Arg(0)
		options.Banner = flag.Arg(1)
	case 3:
		options.Substring = flag.Arg(0)
		options.Input = flag.Arg(1)
		options.Banner = flag.Arg(2)
	}
	return options
}

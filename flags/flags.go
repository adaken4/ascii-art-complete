package flags

import (
	"errors"
	"flag"
)

type Options struct {
	Color     string
	Output    string
	Align     string
	Substring string
	Input     string
	Banner    string
}

func ParseOptions() (Options, error) {
	var options Options
	flag.StringVar(&options.Color, "color", "", "Usage: --color=<color>")
	flag.StringVar(&options.Output, "output", "", "Usage: --ouput=<file.txt>")
	flag.StringVar(&options.Align, "align", "", "Usage: --align=<position>")
	flag.Parse()

	switch len(flag.Args()) {
	case 0:
		return options, errors.New("please provide input text")
	case 1:
		options.Input = flag.Arg(0)
	case 2:
		validBanners := map[string]bool{
			"thinkertoy": true,
			"rounded":    true,
			"shadow":     true,
			"standard":   true,
		}
		if validBanners[flag.Arg(1)] {
			options.Input = flag.Arg(0)
			options.Banner = flag.Arg(1)
		} else {
			options.Substring = flag.Arg(0)
			options.Input = flag.Arg(1)
		}
	case 3:
		options.Substring = flag.Arg(0)
		options.Input = flag.Arg(1)
		options.Banner = flag.Arg(2)
	default:
		return options, errors.New("invalid number of arguments")
	}
	return options, nil
}

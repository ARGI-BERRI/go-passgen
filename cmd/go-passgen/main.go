package main

import (
	"github.com/ARGI_BERRI/go-passgen/internal/passgen"
	"github.com/alexflint/go-arg"
)

func main() {
	var args struct {
		Length  uint8 `arg:"required"`
		Upper   bool
		Lower   bool
		Numeral bool
		Symbol  bool
	}

	arg.MustParse(&args)

	options := passgen.GenerationOption{
		UseUpperCase: args.Upper,
		UseLowerCase: args.Lower,
		UseNumerals:  args.Numeral,
		UseSymbols:   args.Symbol,
	}

	password := passgen.Generate(args.Length, options)
	println(password)
}

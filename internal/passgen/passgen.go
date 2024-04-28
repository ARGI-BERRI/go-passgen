package passgen

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const (
	ALPHABET_UPPER = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ALPHABET_LOWER = "abcdefghijklmnopqrstuvwxyz"
	NUMERALS       = "0123456789"
	SYMBOLS        = "~!@#$%^&*()[]{}<>;:,.|/"
)

type GenerationOption struct {
	UseUpperCase bool
	UseLowerCase bool
	UseNumerals  bool
	UseSymbols   bool
}

func Generate(length uint8, options GenerationOption) string {
	password := make([]string, length)

	charsets, charLength := getUseCharset(options)

	if charLength == 0 {
		return ""
	}

	for i := range password {
		n, err := rand.Int(rand.Reader, big.NewInt(charLength))

		if err != nil {
			panic(err)
		}

		password[i] = string(charsets[n.Int64()])
	}

	return strings.Join(password, "")
}

func getUseCharset(options GenerationOption) (string, int64) {
	charsets := ""
	charLength := 0

	if options.UseUpperCase {
		charsets += ALPHABET_UPPER
		charLength += len(ALPHABET_UPPER)
	}

	if options.UseLowerCase {
		charsets += ALPHABET_LOWER
		charLength += len(ALPHABET_LOWER)
	}

	if options.UseNumerals {
		charsets += NUMERALS
		charLength += len(NUMERALS)
	}

	if options.UseSymbols {
		charsets += SYMBOLS
		charLength += len(SYMBOLS)
	}

	return charsets, int64(charLength)
}

package passgen

import (
	"strings"
	"testing"
)

type TestCase struct {
	name       string
	options    GenerationOption
	charsets   string
	charLength int64
}

func Test_getUseCharset(t *testing.T) {
	for i := range []bool{false} {
		println(i)
	}

	var tests []TestCase

	// TODO: もうちょっとマシな書き方があるはず
	for _, useUpperCase := range []bool{true, false} {
		for _, useLowerCase := range []bool{true, false} {
			for _, useSymbols := range []bool{true, false} {
				for _, useNumerals := range []bool{true, false} {
					var names []string
					charsets := ""
					charLength := 0

					if useUpperCase {
						names = append(names, "uppercase")
						charsets += ALPHABET_UPPER
						charLength += len(ALPHABET_UPPER)
					}

					if useLowerCase {
						names = append(names, "lowercase")
						charsets += ALPHABET_LOWER
						charLength += len(ALPHABET_LOWER)
					}

					if useNumerals {
						names = append(names, "numeral")
						charsets += NUMERALS
						charLength += len(NUMERALS)
					}

					if useSymbols {
						names = append(names, "symbol")
						charsets += SYMBOLS
						charLength += len(SYMBOLS)
					}

					// Append a struct instance to the tests slice
					tests = append(tests, TestCase{
						name: strings.Join(names, ", "),
						options: GenerationOption{
							UseUpperCase: useUpperCase,
							UseLowerCase: useLowerCase,
							UseNumerals:  useNumerals,
							UseSymbols:   useSymbols,
						},
						charsets:   charsets,
						charLength: int64(charLength),
					})
				}
			}
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			charset, charLength := getUseCharset(tt.options)
			if charset != tt.charsets {
				t.Errorf("getUseCharset() charset = %v, charsets %v", charset, tt.charsets)
			}
			if charLength != tt.charLength {
				t.Errorf("getUseCharset() charLength = %v, charsets %v", charLength, tt.charLength)
			}
		})
	}
}

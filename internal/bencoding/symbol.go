package bencoding

import "unicode"

type symbol string

const (
  dictSymbol symbol = "d"
  intSymbol  symbol = "i"
  listSymbol symbol = "l"
  endSymbol  symbol = "e"
)

func (s symbol) isDigit() bool {
  return unicode.IsDigit([]rune(s)[0])
}

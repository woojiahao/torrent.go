package bencoding

import "github.com/woojiahao/torrent.go/internal/utility"

// Parses a string into a TInteger.
func parseTInteger(information string) TInteger {
  result := TIntegerRegex.FindAllStringSubmatch(information, -1)[0]

  data := utility.StrToInt(result[1])

  return TInteger{
    Original: information,
    Data:     data,
  }
}

package bencoding_dep

import (
  "fmt"
  "github.com/woojiahao/torrent.go/internal/utility"
)

// Parses a string into a TString.
func decodeTString(information string) TString {
  result := TStringRegex.FindAllStringSubmatch(information, -1)[0]

  data := result[2]
  length := utility.StrToInt(result[1])

  if length != len(data) {
    panic(fmt.Sprintf("invalid length provided; expected: %d, got: %d", len(data), length))
  }

  return TString{
    Original: information,
    Data:     data,
    Length:   length,
  }
}

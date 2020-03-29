package bencoding

import (
  "fmt"
  "github.com/woojiahao/torrent.go/internal/utility"
  "regexp"
)

var TStringRegex = regexp.MustCompile("^(\\d+):(\\w+)$")
var TIntegerRegex = regexp.MustCompile("^i([-\\d]+)e$")
var TListRegex = regexp.MustCompile("^l(.+)e$")
var TDictionaryRegex = regexp.MustCompile("^(d.+)e$")

// Types starting with 'T' defined as the torrent data types.
type (
  // Base interface for all T-types
  TType interface{}

  // Normal strings [series of continuous characters]
  // Format: <length>:<data>
  // Example: 7:network
  TString struct {
    Original string
    Data     string
    Length   int
  }

  // Normal integers
  // Format: i<integer>e
  // Example: i3e
  TInteger struct {
    Original string
    Data     int
  }

  // List of types [strings, integers, lists, dictionaries]
  // Format: l<contents>e
  // Example: li3ei3ee
  TList struct {
    Original string
    Data     []TType
  }

  // Mapping of keys to values
  // Format: d<keys><values>e
  // Example: d3:onei1e3:twoi2ee
  TDictionary struct {
    Original string
    Data     map[TType]TType
  }
)

func Parse(information string) TType {
  b := []byte(information)
  switch {
  case TStringRegex.Match(b):
    return parseTString(information)
  case TIntegerRegex.Match(b):
    return parseTInteger(information)
  case TListRegex.Match(b):
    panic("not implemented")
  case TDictionaryRegex.Match(b):
    panic("not implemented")
  default:
    panic("invalid information format. please refer to the specification for the appropriate data type format")
  }
}

// Parses a string into a TString.
func parseTString(information string) TString {
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

// Parses a string into a TInteger.
func parseTInteger(information string) TInteger {
  result := TIntegerRegex.FindAllStringSubmatch(information, -1)[0]

  data := utility.StrToInt(result[1])

  return TInteger{
    Original: information,
    Data:     data,
  }
}

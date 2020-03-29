package bencoding

import (
  "regexp"
)

var TStringRegex = regexp.MustCompile("^(\\d+):(\\w+)$")
var TIntegerRegex = regexp.MustCompile("^i([-\\d]+)e$")
var TListRegex = regexp.MustCompile("^l(.+)e$")
var TDictionaryRegex = regexp.MustCompile("^d(.+)e$")

const END = "e"

// Check the README for the bencoding format patterns.

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
    Data     map[string]TType
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
    return parseTDictionary(information)
  default:
    panic("invalid information format. please refer to the specification for the appropriate data type format")
  }
}

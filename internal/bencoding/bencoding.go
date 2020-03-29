package bencoding

import (
  "fmt"
  . "github.com/woojiahao/torrent.go/internal/utility"
  "regexp"
  "strings"
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

func Decode(information string) TType {
  b := []byte(information)
  switch {
  case TStringRegex.Match(b):
    return decodeTString(information)
  case TIntegerRegex.Match(b):
    return decodeTInteger(information)
  case TListRegex.Match(b):
    return decodeTList(information)
  case TDictionaryRegex.Match(b):
    return decodeTDictionary(information)
  default:
    panic("invalid information format. please refer to the specification for the appropriate data type format")
  }
}

func decodeConsecutive(data string, handler func(TType)) {
  counter := 0

  for counter < len(data) {
    cur := string(data[counter])
    jump := 0
    var result TType

    if IsDigit(cur) {
      length := StrToInt(cur)
      content := data[counter : counter+length+2]
      result = decodeTString(content)
      jump = len(content)
    } else if IsStrInRange(cur, "d", "i", "l") {
      // Only assume we can take the first e when the value is an integer
      // If the input is a dictionary or list, the last e should be taken instead
      // TODO Maybe loop endlessly to test if found 'e' is the right e
      jump = strings.Index(data[counter:], END) + 1
      content := data[counter : counter+jump]

      switch cur {
      case "d":
        result = decodeTDictionary(content)
      case "i":
        result = decodeTInteger(content)
      case "l":
        result = decodeTList(content)
      default:
        panic("invalid character")
      }
    } else {
      panic(fmt.Sprintf("invalid character %s", cur))
    }

    counter += jump
    handler(result)
  }
}

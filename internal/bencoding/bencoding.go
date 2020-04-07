package bencoding

import (
  "fmt"
  . "github.com/woojiahao/torrent.go/internal/utility"
  "strings"
  "unicode"
)

type symbol string

const (
  dictSymbol symbol = "d"
  intSymbol  symbol = "i"
  listSymbol symbol = "l"
  endSymbol  symbol = "e"
)

var startSymbols = []symbol{dictSymbol, intSymbol, listSymbol}

func (s symbol) isDigit() bool {
  return unicode.IsDigit([]rune(s)[0])
}

func (s symbol) isStartSymbol() bool {
  for _, sym := range startSymbols {
    if s == sym {
      return true
    }
  }

  return false
}

func Decode(input string) TType {
  result, _, err := decode(input)
  LogCheck(err)
  return result
}

func decode(input string) (result TType, jump int, err error) {
  cur := symbol(input[0])

  if cur.isDigit() {
    result, jump = decodeTString(input)
  } else if cur.isStartSymbol() {
    switch cur {
    case dictSymbol:
      result, jump = decodeTDict(input)
    case intSymbol:
      result, jump = decodeTInt(input)
    case listSymbol:
      result, jump = decodeTList(input)
    }
  } else {
    err = &decodeError{fmt.Sprintf("invalid code: %s", cur)}
  }

  return
}

// Decodes a bencoded string into a TString
func decodeTString(input string) (result TString, jump int) {
  delimiterPos := strings.Index(input, ":")
  length := StrToInt(input[:delimiterPos])
  jump = length + delimiterPos + 1
  result = TString(input[delimiterPos+1 : jump])
  return
}

// Decodes a bencoded string into a TInt
func decodeTInt(input string) (result TInt, jump int) {
  delimiterPos := strings.Index(input, "e")
  result = TInt(StrToInt(input[1:delimiterPos]))
  jump = delimiterPos + 1
  return
}

// Decodes a bencoded string into a TDict
func decodeTDict(input string) (result TDict, jump int) {
  result = make(map[string]TType)

  data, jump := decodeTList(input)

  if len(data)%2 != 0 {
    LogCheck(&decodeError{"invalid dictionary form; make sure that each <key> has a matching <value>"})
  }

  for i := 0; i < len(data); i += 2 {
    key, value := data[i], data[i+1]
    tString, ok := key.(TString)
    if !ok {
      LogCheck(&decodeError{"dictionary key must be a string"})
    }
    result[string(tString)] = value
  }

  return
}

// Decodes a bencoded string into a TList
func decodeTList(input string) (result TList, jump int) {
  data := make([]TType, 0)
  counter := 1
  cur, remaining := string(input[counter]), input[counter:]

  for cur != "e" {
    d, j, err := decode(remaining)

    LogCheck(err)

    data = append(data, d)
    counter += j
    remaining = input[counter:]
    cur = string(input[counter])
  }

  return data, counter + 1
}

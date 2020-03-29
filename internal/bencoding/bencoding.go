package bencoding

import (
  "fmt"
  . "github.com/woojiahao/torrent.go/internal/utility"
  "strings"
)

type (
  TType interface {
    Encode() string
  }

  TString string
  TInt    int
  TList   []TType
  TDict   map[string]TType
)

func (t TString) Encode() string {
  value := string(t)
  return fmt.Sprintf("%d:%s", len(value), value)
}

func (t TInt) Encode() string {
  value := int(t)
  return fmt.Sprintf("i%de", value)
}

func (t TList) Encode() string {
  return "le"
}

func (t TDict) Encode() string {
  return "de"
}

func Decode(input string) TType {
  result, _ := decode(input)
  return result
}

func decode(input string) (result TType, jump int) {
  cur := string(input[0])
  if IsDigit(cur) {
    result, jump = decodeTString(input)
    return
  } else if IsStrInRange(cur, "d", "i", "l") {
    switch cur {
    case "d":
      result, jump = decodeTDict(input)
      return
    case "i":
      result, jump = decodeTInt(input)
      return
    case "l":
      result, jump = decodeTList(input)
    }
  } else {
    panic("invalid code")
  }

  return
}

func decodeTString(input string) (result TString, jump int) {
  delimiterPos := strings.Index(input, ":")
  length := StrToInt(input[:delimiterPos])
  result = TString(input[delimiterPos+1 : length+2])
  jump = length + len(string(length)) + 1
  return
}

func decodeTInt(input string) (result TInt, jump int) {
  delimiterPos := strings.Index(input, "e")
  result = TInt(StrToInt(input[1:delimiterPos]))
  jump = delimiterPos + 1
  return
}

func decodeTDict(input string) (result TDict, jump int) {
  result = make(map[string]TType)

  data, jump := decodeTList(input)

  if len(data)%2 != 0 {
    panic("invalid dictionary form, make sure that there are matching <value>s for each <key>")
  }

  for i := 0; i < len(data); i += 2 {
    key, value := data[i], data[i+1]
    tString, ok := key.(TString)
    if !ok {
      panic("keys must be a string")
    }
    result[string(tString)] = value
  }

  return
}

func decodeTList(input string) (result TList, jump int) {
  data := make([]TType, 0)
  counter := 1
  cur, remaining := string(input[counter]), input[counter:]

  for cur != "e" {
    d, j := decode(remaining)
    data = append(data, d)
    counter += j
    remaining = input[counter:]
    cur = string(input[counter])
  }

  return data, counter + 1
}

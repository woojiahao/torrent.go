package bencoding

import (
  . "github.com/woojiahao/torrent.go/internal/utility"
  "strings"
)

func Decode(input string) TType {
  result, _ := decode(input)
  return result
}

func decode(input string) (result TType, jump int) {
  cur := string(input[0])
  if IsDigit(cur) {
    result, jump = decodeTString(input)
  } else if IsStrInRange(cur, "d", "i", "l") {
    switch cur {
    case "d":
      result, jump = decodeTDict(input)
    case "i":
      result, jump = decodeTInt(input)
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
  jump = length + delimiterPos + 1
  result = TString(input[delimiterPos+1 : jump])
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

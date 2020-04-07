package bencoding

import (
  "fmt"
  . "github.com/woojiahao/torrent.go/internal/utility"
  "strconv"
  "strings"
)

func Decode(input string) TType {
  result, _, err := decode(input)
  Check(err)
  return result
}

func decode(input string) (result TType, jump int, err error) {
  cur := symbol(input[0])

  if cur.isDigit() {
    result, jump, err = decodeTString(input)
    return
  }

  switch cur {
  case dictSymbol:
    result, jump, err = decodeTDict(input)
  case intSymbol:
    result, jump, err = decodeTInt(input)
  case listSymbol:
    result, jump, err = decodeTList(input)
  default:
    err = &decodeError{fmt.Sprintf("invalid code: %s", cur)}
  }

  return
}

func decodeTString(input string) (result TString, jump int, err error) {
  delimiterPos := strings.Index(input, ":")
  if length, err := strconv.Atoi(input[:delimiterPos]); err != nil {
    return
  } else {
    jump := length + delimiterPos + 1
    result = TString(input[delimiterPos+1 : jump])
    return
  }
}

func decodeTInt(input string) (result TInt, jump int, err error) {
  delimiterPos := strings.Index(input, string(endSymbol))
  if i, err := strconv.Atoi(input[1:delimiterPos]); err != nil {
    return
  } else {
    result = TInt(i)
    jump = delimiterPos + 1
    return
  }
}

func decodeTDict(input string) (result TDict, jump int, err error) {
  data, jump, err := decodeContinuous(input)

  if len(data)%2 != 0 {
    err = &decodeError{"invalid dictionary form; make sure that each <key> has a matching <value>"}
    return
  }

  result = make(map[string]TType)
  for i := 0; i < len(data); i += 2 {
    key, value := data[i], data[i+1]

    if tString, ok := key.(TString); !ok {
      err = &decodeError{"dictionary key must be a string"}
      return
    } else {
      result[string(tString)] = value
    }
  }

  return
}

func decodeTList(input string) (TList, int, error) {
  return decodeContinuous(input)
}

// Continually reads and parses the inputs until an end symbol is encountered.
func decodeContinuous(input string) ([]TType, int, error) {
  data := make([]TType, 0)
  counter := 1
  cur, remaining := symbol(input[counter]), input[counter:]

  for cur != endSymbol {
    d, j, err := decode(remaining)
    if err != nil {
      return nil, 0, err
    }

    data = append(data, d)
    counter += j
    remaining = input[counter:]
    cur = symbol(input[counter])
  }

  return data, counter + 1, nil
}

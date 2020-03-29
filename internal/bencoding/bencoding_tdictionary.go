package bencoding

import (
  "fmt"
  . "github.com/woojiahao/torrent.go/internal/utility"
  "strings"
)

// Parses a string into a TDictionary.
// This is broken up as <key1><value1><key2><value2><key3><value3>.
// Keys must be strings and sorted in alphabetical order.
// TODO Check if the keys are in order
func parseTDictionary(information string) TDictionary {
  data := TDictionaryRegex.FindAllStringSubmatch(information, -1)[0][1]

  counter := 0
  isKey, key := true, ""
  dict := make(map[string]TType)

  for counter < len(data) {
    cur := string(data[counter])
    jump := 0

    if IsDigit(cur) {
      length := StrToInt(cur)
      content := data[counter : counter+length+2]
      tString := parseTString(content)
      if isKey {
        key = tString.Data
        dict[key] = nil
      } else {
        dict[key] = tString
      }
      jump = len(content)
    } else if IsStrInRange(cur, "l", "i", "d") {
      if isKey {
        panic("cannot be parsing a key now if the value is not a string")
      } else {
        jump = strings.Index(data[counter:], END) + 1
        content := data[counter : counter+jump]
        var result TType

        switch cur {
        case "i":
          result = parseTInteger(content)
        case "d":
          result = parseTDictionary(content)
        case "l":
          fmt.Println("parsing an list")
        default:
          panic("invalid character")
        }

        dict[key] = result
      }
    } else {
      panic(fmt.Sprintf("invalid character %s", cur))
    }

    counter += jump
    isKey = !isKey
  }

  return TDictionary{
    Original: data,
    Data:     dict,
  }
}

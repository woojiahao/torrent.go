package bencoding

// Parses a string into a TDictionary.
// This is broken up as <key1><value1><key2><value2><key3><value3>.
// Keys must be strings and sorted in alphabetical order.
// TODO Check if the keys are in order
func decodeTDictionary(information string) TDictionary {
  data := TDictionaryRegex.FindAllStringSubmatch(information, -1)[0][1]

  isKey, key := true, ""
  dict := make(map[string]TType)

  decodeConsecutive(
    data,
    func(t TType) {
      if isKey {
        tString, ok := t.(TString)
        if !ok {
          panic("cannot be parsing a key now if the value is not a string.")
        }

        key = tString.Data
        dict[key] = nil
      } else {
        dict[key] = t
      }

      isKey = !isKey
    },
  )

  return TDictionary{
    Original: data,
    Data:     dict,
  }
}

package bencoding_dep

// Parses string to TList.
func decodeTList(information string) TList {
  data := TListRegex.FindAllStringSubmatch(information, -1)[0][1]
  lst := make([]TType, 0)
  decodeConsecutive(data, func(t TType) { lst = append(lst, t) })
  return TList{data, lst}
}

package bencoding

import (
  "fmt"
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

func toTypeStatusCheck(ok bool, t string) {
  if !ok {
    panic(fmt.Sprintf("failed to convert TType to %s", t))
  }
}

func (t TString) Encode() string {
  value := string(t)
  return fmt.Sprintf("%d:%s", len(value), value)
}

func ToString(t TType) TString {
  if t == nil {
    return ""
  }
  tString, ok := t.(TString)
  toTypeStatusCheck(ok, "TString")
  return tString
}

func (t TString) Value() string {
  return string(t)
}

func (t TInt) Encode() string {
  value := int(t)
  return fmt.Sprintf("i%de", value)
}

func ToInt(t TType) TInt {
  if t == nil {
    return 0
  }
  tInt, ok := t.(TInt)
  toTypeStatusCheck(ok, "TInt")
  return tInt
}

func (t TInt) Value() int {
  return int(t)
}

// TODO Add encoding to this
func (t TList) Encode() string {
  return "le"
}

func ToList(t TType) TList {
  if t == nil {
    return nil
  }
  tList, ok := t.(TList)
  toTypeStatusCheck(ok, "TList")
  return tList
}

func (t TList) Value() []TType {
  return t
}

// TODO Add encoding to this
func (t TDict) Encode() string {
  return "de"
}

func ToDict(t TType) TDict {
  if t == nil {
    return nil
  }
  tDict, ok := t.(TDict)
  toTypeStatusCheck(ok, "TDict")
  return tDict
}

func (t TDict) Value() map[string]TType {
  return t
}

func (t TDict) String() string {
  data := make([]string, 0)
  for k, v := range t {
    data = append(data, fmt.Sprintf("%s: %s", k, v))
  }
  return strings.Join(data, "\n")
}
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

func ToString(t TType) TString {
  if t == nil {
    return ""
  }

  tString, ok := t.(TString)
  Check(toTypeStatusCheck(ok, "TString"))

  return tString
}

func (t TString) Value() string {
  return string(t)
}

func (t TInt) Encode() string {
  value := int(t)
  return fmt.Sprintf("%s%d%s", intSymbol, value, endSymbol)
}

func ToInt(t TType) TInt {
  if t == nil {
    return 0
  }

  tInt, ok := t.(TInt)
  Check(toTypeStatusCheck(ok, "TInt"))

  return tInt
}

func (t TInt) Value() int {
  return int(t)
}

func (t TList) Encode() string {
  values := make([]string, 0)
  for _, i := range t {
    values = append(values, i.Encode())
  }

  return fmt.Sprintf("%s%s%s", listSymbol, strings.Join(values, ""), endSymbol)
}

func ToList(t TType) TList {
  if t == nil {
    return nil
  }

  tList, ok := t.(TList)
  Check(toTypeStatusCheck(ok, "TList"))

  return tList
}

func (t TList) Value() []TType {
  return t
}

func (t TDict) Encode() string {
  values := make([]string, 0)
  for key, value := range t {
    values = append(values, fmt.Sprintf("%s%s", TString(key).Encode(), value.Encode()))
  }

  return fmt.Sprintf("%s%s%s", dictSymbol, strings.Join(values, ""), endSymbol)
}

func ToDict(t TType) TDict {
  if t == nil {
    return nil
  }

  tDict, ok := t.(TDict)
  Check(toTypeStatusCheck(ok, "TDict"))

  return tDict
}

func (t TDict) Value() map[string]TType {
  return t
}

func toTypeStatusCheck(ok bool, t string) error {
  if !ok {
    return &typeConversionError{t}
  }

  return nil
}

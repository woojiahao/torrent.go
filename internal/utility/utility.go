package utility

import (
  "strconv"
  "unicode"
)

func Check(err error) {
  if err != nil {
    panic(err)
  }
}

func StrToInt(in string) int {
  val, err := strconv.Atoi(in)
  Check(err)
  return val
}

func StrToRune(in string) rune {
  return []rune(in)[0]
}

func IsDigit(in string) bool {
  return unicode.IsDigit(StrToRune(in))
}

func IsStrInRange(in string, ch ...string) bool {
  for _, c := range ch {
    if c == in {
      return true
    }
  }

  return false
}

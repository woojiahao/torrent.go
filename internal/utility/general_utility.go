package utility

import (
  "math/rand"
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

func randomInt(min, max int) int {
  return min + rand.Intn(max-min)
}

func randomChar() byte {
  isCapital := randomInt(0, 1)
  switch isCapital {
  case 0:
    return byte(randomInt(97, 122))
  case 1:
    return byte(randomInt(65, 90))
  default:
    panic("invalid int")
  }
}

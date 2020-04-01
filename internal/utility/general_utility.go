package utility

import (
  "math/rand"
  "strconv"
  "unicode"
)

// Checks if an error is non-nil; if non-nil, panic with the error; else ignore
func Check(err error) {
  if err != nil {
    panic(err)
  }
}

// Converts a string to an integer
func StrToInt(in string) int {
  val, err := strconv.Atoi(in)
  Check(err)
  return val
}

// Converts a string to a rune
func StrToRune(in string) rune {
  return []rune(in)[0]
}

// Check if a string is a digit
func IsDigit(in string) bool {
  return unicode.IsDigit(StrToRune(in))
}

// Check if a string is within the range of specified strings
func IsStrInRange(in string, ch ...string) bool {
  for _, c := range ch {
    if c == in {
      return true
    }
  }

  return false
}

// Generates a random integer from min (inclusive) to max (exclusive)
func RandomInt(min, max int) int {
  return min + rand.Intn(max-min)
}

// Generates a random character - lowercase and uppercase
func RandomChar() byte {
  isCapital := RandomInt(0, 2)
  switch isCapital {
  case 0:
    return byte(RandomInt(97, 123))
  case 1:
    return byte(RandomInt(65, 91))
  default:
    panic("invalid int")
  }
}

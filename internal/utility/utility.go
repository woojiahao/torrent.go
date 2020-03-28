package utility

import "strconv"

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

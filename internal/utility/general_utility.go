package utility

import (
  "crypto/sha1"
  "encoding/binary"
  "hash"
  "log"
  "math"
  "math/rand"
  "strconv"
)

// Checks if an error is non-nil; if non-nil, panic with the error; else ignore
func Check(err error) {
  if err != nil {
    panic(err)
  }
}

func LogCheck(err error) {
  if err != nil {
    log.Fatal(err.Error())
  }
}

// Converts a string to an integer
func StrToInt(in string) int {
  val, err := strconv.Atoi(in)
  LogCheck(err)
  return val
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

// Generate a SHA1 Hash for a given string input
func GenerateSHA1Hash(input string) hash.Hash {
  h := sha1.New()
  h.Write([]byte(input))
  return h
}

func ToBigEndian(value, size int) []byte {
  buf := make([]byte, size)
  binary.BigEndian.PutUint32(buf, uint32(value))
  return buf
}

func FromBigEndian(value []byte) int {
  return int(binary.BigEndian.Uint32(value))
}

// Implementation of math.Pow for int
func Pow(base, pow int) int {
  return int(math.Pow(float64(base), float64(pow)))
}

// Implementation of math.Min for int
func Min(a, b int) int {
  return int(math.Min(float64(a), float64(b)))
}

package utility

import (
  "crypto/sha1"
  "encoding/binary"
  "hash"
  "log"
  "math"
  "math/rand"
)

func Check(err error) {
  if err != nil {
    log.Fatal(err.Error())
  }
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
// The info_hash is the SHA1 hash representation of the bencoding info portion of the metadata
// The SHA1 hash generated is 40 characters long for human reading, it is in fact a hex string
// The tracker must receive the URL-encoded version of the hex string
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

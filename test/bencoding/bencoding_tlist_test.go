package bencoding

import (
  "fmt"
  "github.com/woojiahao/torrent.go/internal/bencoding"
  "testing"
)

func TestTListDecodeValidFormat(t *testing.T) {
  result := bencoding.Decode("l3:onei8ed3:twoi2eee")
  fmt.Println(result)
}

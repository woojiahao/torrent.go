package bencoding

import (
  . "github.com/woojiahao/torrent.go/internal/bencoding"
  "testing"
)

func TestTIntegerParseValidFormat(t *testing.T) {
  data := []TInteger{
    {"i3e", 3},
    {"i0e", 0},
    {"i-1e", -1},
    {"i99999e", 99999},
  }

  for _, d := range data {
    result := Parse(d.Original)

    tInteger, ok := result.(TInteger)
    if !ok {
      t.Errorf("%s was converted to a %T, instead of TInteger", d.Original, result)
    }

    if d != tInteger {
      t.Errorf("%s was not parsed correctly; expected: %d, got: %d", d.Original, d.Data, tInteger.Data)
    }

    t.Logf("string successfully converted to TInteger: %v", tInteger)
  }
}

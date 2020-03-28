package bencoding

import (
  "github.com/woojiahao/torrent.go/internal/bencoding"
  "testing"
)

func TestTStringParseValidFormat(t *testing.T) {
  data := []bencoding.TString{
    {"4:test", "test", 4},
    {"7:network", "network", 7},
    {"1:i", "i", 1},
    {"2:hi", "hi", 2},
  }

  for _, d := range data {
    result := bencoding.Parse(d.Original)
    tString, ok := result.(bencoding.TString)
    if !ok {
      t.Errorf("%s was converted to a %T, instead of a TString", d.Original, result)
    }

    if d != result {
      t.Errorf(
        "%s was not parsed correctly; expected: %s of %d characters, got: %s of %d characters",
        d.Original,
        d.Data,
        d.Length,
        tString.Data,
        tString.Length,
      )
    }

    t.Logf("string successfully converted to TString: %v", tString)
  }
}

func TestTStringParseInvalidFormat(t *testing.T) {

}

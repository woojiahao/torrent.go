package bencoding

import (
  "testing"
)

func TestTStringParseValidFormat(t *testing.T) {
  //data := []bencoding_dep.TString{
  //  {"4:test", "test", 4},
  //  {"7:network", "network", 7},
  //  {"1:i", "i", 1},
  //  {"2:hi", "hi", 2},
  //}
  //
  //for _, d := range data {
  //  result := bencoding_dep.Decode(d.Original)
  //  tString, ok := result.(bencoding_dep.TString)
  //  if !ok {
  //    t.Errorf("%s was converted to a %T, instead of a TString", d.Original, result)
  //  }
  //
  //  if d != tString {
  //    t.Errorf(
  //      "%s was not parsed correctly; expected: %s of %d characters, got: %s of %d characters",
  //      d.Original,
  //      d.Data,
  //      d.Length,
  //      tString.Data,
  //      tString.Length,
  //    )
  //  }
  //
  //  t.Logf("string successfully converted to TString: %v", tString)
  //}
}

func TestTStringParseInvalidFormat(t *testing.T) {
  //defer func() {
  //  if r := recover(); r == nil {
  //    t.Errorf("the code did not panic")
  //  }
  //}()
  //
  //data := []bencoding_dep.TString{
  //  {"5:test", "test", 5},
  //  {"7:network", "network", 8},
  //  {"1:i", "", 1},
  //  {"2:h", "hi", 2},
  //}
  //
  //for _, d := range data {
  //  result := bencoding_dep.Decode(d.Original)
  //  tString := result.(bencoding_dep.TString)
  //
  //  if tString == d {
  //    t.Errorf("%s should be incorrectly parsed", tString.Original)
  //  }
  //}
}

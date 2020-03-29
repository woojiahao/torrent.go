package bencoding

import (
  "fmt"
  . "github.com/woojiahao/torrent.go/internal/bencoding"
  "testing"
)

func TestTDictionaryParseValidFormat(t *testing.T) {
  dictionary := make(map[string]TType)
  dictionary["one"] = TInteger{
    Original: "i1e",
    Data:     1,
  }
  dictionary["str"] = TString{
    Original: "3:str",
    Data:     "str",
    Length:   3,
  }
  data := Decode("d3:onei1e3:str3:stre")
  fmt.Println(data)
  //data := []TDictionary{
  //  {"d3:onei1ee", map[string]TType{
  //    "one": TInteger{
  //      Original: "i1e",
  //      Data:     1,
  //    },
  //  },}
  //}
}

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
  dictionary["dict"] = TDictionary{
    Original: "d3:onei1ee",
    Data: map[string]TType{
      "one": TInteger{
        Original: "i1e",
        Data:     1,
      },
    },
  }
  data := Decode("d3:onei1e3:str3:str3:keyl3:onei8ed3:twoi2eeee")
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

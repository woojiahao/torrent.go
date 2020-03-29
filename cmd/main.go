package main

import (
  "fmt"
  "github.com/woojiahao/torrent.go/internal/bencoding"
)

func main() {
  foo := bencoding.Decode("3:Onei31e")
  bar := bencoding.Decode("2:TE")
  baz := bencoding.Decode("d3:onei1e5:threei3e3:lstd3:onei1eee")
  lst := bencoding.Decode("li1e3:onel3:oneed3:one3:oneee")
  fmt.Println(foo, bar)
  fmt.Println(baz, lst)
}

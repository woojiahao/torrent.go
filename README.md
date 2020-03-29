![.github/workflows/main.yml](https://github.com/woojiahao/torrent.go/workflows/.github/workflows/main.yml/badge.svg?branch=master)

# torrent.go
Implementation of the BitTorrent protocol in Golang

## References

### BitTorrent specification

- [How does BitTorrent work - Overview](https://www.howtogeek.com/141257/htg-explains-how-does-bittorrent-work/)
- [Understanding the BitTorrent specification](http://dandylife.net/docs/BitTorrent-Protocol.pdf)
- [Implementing BitTorrent using AsyncIO in Python](https://youtu.be/Pe3b9bdRtiE)

### Bencoding format

- [Wikipedia](https://en.wikipedia.org/wiki/Bencode)
- [BitTorrent specification](https://www.bittorrent.org/beps/bep_0003.html)

## TODO 

- [ ] Set the valid parsing of the bencoding to use defers to catch the appropriate panics
- [ ] Add data validation for bencoding decoding
- [ ] Add encoding to bencoding format
- [ ] Update test suite and any other location where "parse" is used instead of decode

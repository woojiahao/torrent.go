![.github/workflows/main.yml](https://github.com/woojiahao/torrent.go/workflows/.github/workflows/main.yml/badge.svg?branch=master)

# torrent.go
Implementation of the BitTorrent protocol in Golang

## Sample file formats

### Single file

```json
{
   "announce": "https://torrent.ubuntu.com/announce",
   "comment": "Ubuntu-MATE CD cdimage.ubuntu.com",
   "creation date": 1581513754,
   "info": {
      "length": 1964244992,
      "name": "ubuntu-mate-18.04.4-desktop-amd64.iso",
      "piece length": 524288,
      "pieces": "..."
   }
}
```

### Multiple files

```json
{
   "created by": "kimbatt.github.io/torrent-creator",
   "creation date": 1585615864,
   "info": {
      "files": [
         {
            "length": 324746,
            "path": [
               "6VEmY3n.jpg"
            ]
         },
         {
            "length": 80702,
            "path": [
               "203wnl7t91321.jpg"
            ]
         },
         {
            "length": 331746,
            "path": [
               "5xzu25hn1lm21.png"
            ]
         }
      ],
      "length": 0,
      "name": "test",
      "piece length": 16384,
      "pieces": "..."
   }
}
```

## Note

For `.torrent` files downloaded from the Internet Archive, the client requires web seeding support. This is clearly indicated 
in the torrent files downloaded. This has yet to be added to `torrent.go` so there might be a limitation to the features of 
the system.

## References

### BitTorrent specification

- [How does BitTorrent work - Overview](https://www.howtogeek.com/141257/htg-explains-how-does-bittorrent-work/)
- [Understanding the BitTorrent specification](http://dandylife.net/docs/BitTorrent-Protocol.pdf)
- [Implementing BitTorrent using AsyncIO in Python](https://youtu.be/Pe3b9bdRtiE)
- [Torrent files](https://en.wikipedia.org/wiki/Torrent_file)
- [Specification](https://wiki.theory.org/index.php/BitTorrentSpecification)

### Bencoding format

- [Wikipedia](https://en.wikipedia.org/wiki/Bencode)
- [BitTorrent specification](https://www.bittorrent.org/beps/bep_0003.html)

### Sample torrents

- [WebTorrent free torrents](https://webtorrent.io/free-torrents)
- [Internet archive](https://archive.org/)

## TODO 

- [ ] Set the valid parsing of the bencoding to use defers to catch the appropriate panics
- [ ] Add data validation for bencoding decoding
- [X] Add encoding to bencoding format
- [ ] Update test suite and any other location where "parse" is used instead of decode
- [ ] Retry downloading if the original torrent failed
- [ ] Support web seeding

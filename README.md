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

## Peer protocol process

The peer protocol is performed over a TCP connection. As there may be multiple peers to be connected to at once, each peer 
connection will maintain its own flow.

The key actors of the process is the **client** and **peer**.

1. Client attempts to establish a TCP connection to the peer
2. If the peer is available, the peer will allow the TCP connection
3. Client sends the a handshake request to the peer with the following format: `<pstrlen><pstr><reserved><info_hash><peer_id>`
4. If the peer wants to communicate with the client, it will respond with a response using the same format; this will mean that
    the TCP connection has been established and the client can begin requesting the peer for pieces of the file
5. Client and peer will communciate in an alternating fashion
6. If the client already has some pieces of the file, it can send a `bitfield` request to the peer; this request informs the
    peer of the bits that it already has
    
Ideally, the moment the client establishes the connection with the peer, it will inform the peer that it is interested and 
is not choked. Then, begin to request for the pieces of data. Once a piece is received, the client will use the pieces bytes
from the torrent file to verify if the received piece has a SHA1 hash that is matches the `.torrent` file. Once verified, the
client can inform the peer that it has the piece via a `have` request. Repeat the request process while the client remains 
interested and the peer remains unchoked.

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
- [ ] Set up event bus system for when adding CLI 
- [ ] Add proper logging and error handling to allow for dynamic system
- [ ] Clean up bencoding code (e.g. place the counter variable in the for loop)
- [ ] Flatten out the folder structure to allow the torrent related files to be split out respectively
- [ ] Use `select{}` with `time.After` to toggle a retry mechanism for server retries

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

### Post-investigation

After some investigation and grokking, I had come to the realisation that the original idea of how peers communicated was naive.
I will be focusing on the P2P aspect of the communication, assuming the handshake and connection has already been established.

To split the pieces, we need to think of it as such. A file is split into pieces. We know the order of these pieces because we 
have the metainfo of the torrent file. This gives us two pieces of information - piece position and SHA1 hash. We can also deduce
how long the piece will be as the metainfo announces to us the `piece length`. We know that we have to query each piece individually,
so we want to split this pieces array out and create a piece of work to perform. This piece of work will contain the index of the 
piece in the file, the hash of the piece, and the length of it. The reason why we need to add the length is because the last piece in
the file might be slightly less than the others. So it's best to include it. 

Then we know that each piece is split into blocks. According to the documentation, current implementations of block sizes usually
have the size to be 16kiB. This also means that we have to keep track of which blocks we receive. This is easier because we can do 
this sequentially. To track the piece progress, we need several things, we need the piece we are tracking - or the index of the 
piece, the current client connection, how much is downloaded, the data that is downloaded, and finally what we are requesting. To
keep track of what we have downloaded, we continually request for blocks with the fixed block size. As we are downloading these 
bytes, we will add to the downloaded section. Ideally, by the end of the download, the downloaded should be the same length as the
piece length. The requsted size is the same as the downloaded size and we must increment both. The only difference is that the 
requested size increases after each `REQUEST` while the downloaded size increases after each confirmed `PIECE`.

Once we download the entire piece, we then check the piece against the SHA1 hash. Once validated, we can send a `HAVE` to the peer
to confirm the download.

```go
package pieces

type pieceWork struct {
  index  int
  hash   [20]byte
  length int
}
```


To implement the logic in goroutines, we fill up a channel with the pieces we want to retrieve. Then, we have another channel to
receive the downloaded pieces. For the work queue, we can actually just iterate over the channel until it is empty. For the 
downloaded pieces, we need the index of the piece and the bytes of the piece:

```go
package pieces

type piece struct {
  index int
  buf   []byte
}
```

Once the handshake is completed, the peer we are connecting to must send us a **bitfield**. The definition of a bitfield was 
originally quite confusing. While it was clear that it represents the pieces present on a peer, it doesn't specify who should be
the one emitting the initial bitfield or how it interacts with the algorithm. Turns out, the bitfield is used to represent what
pieces a peer has. This means that in order to understand what pieces I can request from the peer, the client must keep track
of what pieces the peer has offered. This will represent which pieces we can request from the peer. BUT it is also important to 
know that the peer might not advertise all pieces at once. As such, it can also use `HAVE` messages to inform the client that new
pieces are available.

## References

### BitTorrent specification

- [How does BitTorrent work - Overview](https://www.howtogeek.com/141257/htg-explains-how-does-bittorrent-work/)
- [Understanding the BitTorrent specification](http://dandylife.net/docs/BitTorrent-Protocol.pdf)
- [Implementing BitTorrent using AsyncIO in Python](https://youtu.be/Pe3b9bdRtiE)
- [Torrent files](https://en.wikipedia.org/wiki/Torrent_file)
- [Specification](https://wiki.theory.org/index.php/BitTorrentSpecification)
- [How to write a BitTorrent client](http://www.kristenwidman.com/blog/71/how-to-write-a-bittorrent-client-part-2/)

### Bencoding format

- [Wikipedia](https://en.wikipedia.org/wiki/Bencode)
- [BitTorrent specification](https://www.bittorrent.org/beps/bep_0003.html)

### Sample torrents

- [WebTorrent free torrents](https://webtorrent.io/free-torrents)
- [Internet archive](https://archive.org/)

## TODO 

- [X] ~~Set the valid parsing of the bencoding to use defers to catch the appropriate panics~~
- [X] Add data validation for bencoding decoding
- [X] Add encoding to bencoding format
- [ ] Update test suite and any other location where "parse" is used instead of decode
- [ ] Retry downloading if the original torrent failed
- [ ] Support web seeding
- [ ] Set up event bus system for when adding CLI 
- [ ] Add proper logging and error handling to allow for dynamic system
- [X] Clean up bencoding code (e.g. place the counter variable in the for loop)
- [X] Flatten out the folder structure to allow the torrent related files to be split out respectively
- [ ] Use `select{}` with `time.After` to toggle a retry mechanism for server retries

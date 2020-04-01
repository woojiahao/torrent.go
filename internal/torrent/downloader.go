package torrent

import (
  "fmt"
  . "github.com/woojiahao/torrent.go/internal/utility"
)

// Initiates the TCP connection to begin downloading the torrent information.
// Not to be confused with torrent.Download which is used to initialize the entire torrent
// downloading process.
func download(response *trackerResponse) {
  for _, peer := range response.peers {
    conn := TCP(peer.address(), 3)
    fmt.Println(conn)
  }
}

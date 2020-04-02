package torrent

import (
  . "github.com/woojiahao/torrent.go/internal/utility"
  "log"
)

type handshake struct {
  pstrlen
}

func connectPeer(address string) {
  log.Print("connecting to ", address)
  conn, err := TCP(address, 3)
  if err != nil {
    log.Printf("connection to %s rejected", address)
    return
  }

}

// Initiates the TCP connection to begin downloading the torrent information.
// Not to be confused with torrent.Download which is used to initialize the entire torrent
// downloading process.
// When establishing the TCP connections, we should be ignoring the ones that refuse the connection
// or timeout the request.
func connectToPeers(response *trackerResponse) {
  for _, peer := range response.peers {
    address := peer.address()
    log.Print("connecting to ", address)
    conn, err := TCP(address, 3)
    if err != nil {
      log.Printf("connection to %s rejected", address)
      continue
    }
  }
}

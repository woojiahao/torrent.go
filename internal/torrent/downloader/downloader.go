package downloader

import (
  "github.com/woojiahao/torrent.go/internal/torrent/downloader/handshake"
  "github.com/woojiahao/torrent.go/internal/torrent/tracker"
  . "github.com/woojiahao/torrent.go/internal/utility"
  "log"
)

// Establishes a connection with a peer and continues to use the peer to
// obtain pieces until the peer has opted out of the process
func connectPeer(address string, h *handshake.Handshake) {
  log.Print("connecting to ", address)
  conn, err := TCP(address, 3)
  if err != nil {
    log.Printf("connection to %s rejected", address)
    return
  }

  err = handshake.Request(conn, h)
  if err != nil {
    log.Printf("connection to %s dropped due to reason %s", address, err)
    return
  }

  choked, interested := 2, 0
}

// Initiates the TCP connection to begin downloading the torrent information.
// Not to be confused with torrent.Download which is used to initialize the entire torrent
// downloading process.
// When establishing the TCP connections, we should be ignoring the ones that refuse the connection
// or timeout the request.
func Download(peers []tracker.Peer, infoHash, peerID string) {
  h := handshake.New(infoHash, peerID)
  for _, peer := range peers {
    connectPeer(peer.Address(), h)
  }
}

package downloader

import (
  . "github.com/woojiahao/torrent.go/internal/connection"
  . "github.com/woojiahao/torrent.go/internal/p2p"
  "github.com/woojiahao/torrent.go/internal/tracker"
  "log"
)

type ClientState struct {
  isChoked     bool
  isInterested bool
}

// Establishes a connection with a peer and continues to use the peer to
// obtain pieces until the peer has opted out of the process
func connectPeer(address string, h *Handshake) {
  log.Println("connecting to", address)
  conn, err := TCP(address, 3)
  log.Println("connecting to ", address)
  if err != nil {
    log.Printf("connection to %s rejected", address)
    return
  }

  err = Request(conn, h)
  if err != nil {
    log.Printf("connection to %s dropped due to reason %s", address, err)
    return
  }

  log.Printf("connected to %s", address)

  //clientState := ClientState{false, true}

  err = StartDownloadWorker(conn)
  if err != nil {
    log.Print(err.Error())
    return
  }
}

// Initiates the TCP connection to begin downloading the torrent information.
// Not to be confused with torrent.Download which is used to initialize the entire torrent
// downloading process.
// When establishing the TCP connections, we should be ignoring the ones that refuse the connection
// or timeout the request.
func Download(peers []tracker.Peer, infoHash, peerID string) {
  h := New(infoHash, peerID)
  for _, peer := range peers {
    connectPeer(peer.Address(), h)
  }
}

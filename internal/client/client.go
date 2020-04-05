package client

import (
  . "github.com/woojiahao/torrent.go/internal/bitfield"
  . "github.com/woojiahao/torrent.go/internal/connection"
  . "github.com/woojiahao/torrent.go/internal/tracker"
)

// Client connection to a peer
type Client struct {
  Conn Connection
  Peer
  Choked bool
  Bitfield
  InfoHash [20]byte
  PeerID   [20]byte
}

package client

import (
  "errors"
  "fmt"
  . "github.com/woojiahao/torrent.go/internal/bitfield"
  . "github.com/woojiahao/torrent.go/internal/connection"
  "github.com/woojiahao/torrent.go/internal/handshake"
  "github.com/woojiahao/torrent.go/internal/message"
  . "github.com/woojiahao/torrent.go/internal/tracker"
  "log"
)

// Client connection to a peer
type Client struct {
  Conn *Connection
  Peer
  Choked     bool
  Interested bool
  Bitfield
  InfoHash string
  PeerID   string
}

// Creates a new client connection
func New(peer Peer, infoHash, peerID string) (*Client, error) {
  address := peer.Address()
  log.Println("connecting to", address)

  conn, err := TCP(address, 3)
  if err != nil {
    return nil, errors.New(fmt.Sprintf("connection to %s rejected", address))
  }

  h := handshake.New(infoHash, peerID)
  err = handshake.Request(conn, h)
  if err != nil {
    return nil, errors.New(fmt.Sprintf("connection to %s dropped due to reason %s", address, err))
  }

  log.Println("connected to", address)

  log.Println("requesting for bitfield")

  bitfield, err := receiveBitfield(conn)
  if err != nil {
    return nil, err
  }

  return &Client{
    conn,
    peer,
    true,
    false,
    bitfield,
    infoHash,
    peerID,
  }, nil
}

func receiveBitfield(conn *Connection) (Bitfield, error) {
  m, err := conn.Receive()
  if err != nil {
    return nil, err
  }

  if m.MessageID != message.BitfieldID {
    return nil, errors.New("invalid message type, peer must provide bitfield to provide client with available pieces")
  }

  return m.Payload, nil
}

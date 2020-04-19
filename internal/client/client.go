package client

import (
  "errors"
  "fmt"
  . "github.com/woojiahao/torrent.go/internal/bitfield"
  . "github.com/woojiahao/torrent.go/internal/connection"
  "github.com/woojiahao/torrent.go/internal/handshake"
  "github.com/woojiahao/torrent.go/internal/message"
  . "github.com/woojiahao/torrent.go/internal/tracker"
  . "github.com/woojiahao/torrent.go/internal/utility"
  "log"
)

// Client connection to a peer
// Choked refers to whether the peer is choked
type Client struct {
  Conn *Connection
  Peer
  Choked bool
  Bitfield
  InfoHash string
  PeerID   string
}

func receiveBitfield(conn *Connection) (Bitfield, error) {
  msg, err := message.Read(conn)
  if err != nil {
    return nil, err
  }

  if msg.MessageID != message.BitfieldID {
    return nil, errors.New("invalid message type, peer must provide bitfield to provide client with available pieces")
  }

  return msg.Payload, nil
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
    bitfield,
    infoHash,
    peerID,
  }, nil
}

func (c *Client) SendUnchoke() error {
  unchoke := message.New(message.UnchokeID)
  err := c.Conn.Send(unchoke.Serialize())
  return err
}

func (c *Client) SendInterested() error {
  interested := message.New(message.InterestedID)
  err := c.Conn.Send(interested.Serialize())
  return err
}

func (c *Client) SendRequest(index, begin, length int) error {
  payload := make([]byte, 0, 12)
  payload = append(payload, ToBigEndian(index, 4)...)
  payload = append(payload, ToBigEndian(begin, 4)...)
  payload = append(payload, ToBigEndian(length, 4)...)
  request := message.New(message.RequestID, payload...)
  err := c.Conn.Send(request.Serialize())
  return err
}

func (c *Client) SendHave(index int) error {
  have := message.New(message.HaveID, ToBigEndian(index, 4)...)
  err := c.Conn.Send(have.Serialize())
  return err
}

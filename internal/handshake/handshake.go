package handshake

import (
  "errors"
  "fmt"
  . "github.com/woojiahao/torrent.go/internal/connection"
  "log"
)

const pstr = "BitTorrent protocol"

type Handshake struct {
  pstrlen  int
  pstr     string
  reserved [8]byte
  infoHash string
  peerID   string
}

// Serialize a handshake into bytes to be sent to the TCP server
func (h *Handshake) serialize() []byte {
  buf := make([]byte, h.pstrlen+49)
  buf[0] = byte(h.pstrlen)
  cur := 1
  cur += copy(buf[cur:], h.pstr)
  cur += copy(buf[cur:], h.reserved[:])
  cur += copy(buf[cur:], h.infoHash[:])
  cur += copy(buf[cur:], h.peerID[:])
  return buf
}

// Deserialize an array of bytes received from a TCP server
// The bytes received are in the same format as the ones that are serialized
func deserialize(b []byte) *Handshake {
  pstrLen := int(b[0])
  cur := 1
  pstr, cur := string(b[cur:cur+pstrLen]), cur+pstrLen
  buf, cur := b[cur:cur+8], cur+8
  var reserved [8]byte
  copy(reserved[:], buf[:])
  infoHash, cur := string(b[cur:cur+20]), cur+20
  peerID := string(b[cur:])

  return &Handshake{
    pstrLen,
    pstr,
    reserved,
    infoHash,
    peerID,
  }
}

func New(infoHash, peerID string) *Handshake {
  return &Handshake{
    len(pstr),
    pstr,
    [8]byte{},
    infoHash,
    peerID,
  }
}

func Request(conn *Connection, h *Handshake) error {
  log.Print("requesting handshake")
  err := conn.Send(h.serialize())
  if err != nil {
    log.Fatalf("error occured %s", err.Error())
  }

  buf, err := conn.Receive(h.pstrlen+49)
  if err != nil {
    return err
  }

  if len(buf) == 0 {
    return fmt.Errorf("empty buffer received")
  }

  response := deserialize(buf)

  if response.infoHash != h.infoHash {
    return errors.New(
      fmt.Sprintf(
        "info hash returned by peer does not match client's info hash; expected: %v, given: %v",
        h.infoHash,
        response.infoHash,
      ),
    )
  }

  return nil
}

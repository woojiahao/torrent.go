package torrent

import (
  "errors"
  . "fmt"
  . "github.com/woojiahao/torrent.go/internal/utility"
  "log"
  "net"
)

const pstr = "BitTorrent protocol"

type handshake struct {
  pstrlen  int
  pstr     string
  reserved [8]byte
  infoHash string
  peerID   string
}

func (h *handshake) Serialize() []byte {
  buf := make([]byte, h.pstrlen+49)
  buf[0] = byte(h.pstrlen)
  cur := 1
  cur += copy(buf[cur:], h.pstr)
  cur += copy(buf[cur:], h.reserved[:])
  cur += copy(buf[cur:], h.infoHash[:])
  cur += copy(buf[cur:], h.peerID[:])
  return buf
}

func buildHandshakeRequest(infoHash, peerID string) *handshake {
  return &handshake{
    len(pstr),
    pstr,
    [8]byte{},
    infoHash,
    peerID,
  }
}

func connectPeer(address string, handshake *handshake) {
  log.Print("connecting to ", address)
  conn, err := TCP(address, 3)
  if err != nil {
    log.Printf("connection to %s rejected", address)
    return
  }

  err = handshakeRequest(conn, handshake)
  if err != nil {
    log.Printf("connection to %s dropped due to reason %s", address, err)
  }
}

func handshakeRequest(conn net.Conn, h *handshake) error {
  defer func() {
    _ = conn.Close()
  }()
  // conn.Write returns an int specifying the length of the message
  _, err := conn.Write(h.Serialize())
  if err != nil {
    log.Fatalf("error occured %s", err.Error())
  }

  buf := make([]byte, h.pstrlen+49)
  _, err = conn.Read(buf)
  if err != nil {
    return err
  }
  response := string(buf)
  responseInfoHash := response[28:len(response)-20]
  if responseInfoHash != h.infoHash {
    return errors.New(
      Sprintf(
        "info hash returned by peer does not match client's info hash; expected: %v, given: %v",
        h.infoHash,
        responseInfoHash,
      ),
    )
  }

  return nil
}

// Initiates the TCP connection to begin downloading the torrent information.
// Not to be confused with torrent.Download which is used to initialize the entire torrent
// downloading process.
// When establishing the TCP connections, we should be ignoring the ones that refuse the connection
// or timeout the request.
func download(peers []peer, infoHash, peerID string) {
  handshake := buildHandshakeRequest(infoHash, peerID)
  for _, peer := range peers {
    connectPeer(peer.address(), handshake)
  }
}

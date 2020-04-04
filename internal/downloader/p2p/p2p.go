package p2p

import (
  "errors"
  "fmt"
  "github.com/woojiahao/torrent.go/internal/downloader"
  . "github.com/woojiahao/torrent.go/internal/utility"
)

const maxBufferSize = 16384

// The peer protocol is an alternating stream of length prefixes and messages
func StartDownloadWorker(conn *TCPConn, clientState *downloader.ClientState) error {
  defer func() {
    _ = conn.Conn.Close()
  }()

  for {
    err := conn.SendMessage(UnchokeID)
    err = conn.SendMessage(InterestedID)

    err = conn.SendMessage(RequestID)

    // TODO Alter the byte length to be smaller if large bytes are not needed
    buf, err := conn.Receive(maxBufferSize)
    if err != nil {
      // TODO factor in EOF errors and ignore those
      return errors.New(fmt.Sprintf("connection encountered error %s", err.Error()))
    }
    msg := Deserialize(buf)

    fmt.Println(msg)
  }
}

// Read the server's messages
func readMessage(m *Message) {

}

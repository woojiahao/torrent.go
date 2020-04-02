package p2p

import (
  "errors"
  "fmt"
  "net"
)

func Peer2Peer(conn net.Conn) error {
  for {
    // TODO Alter the byte length to be smaller if large bytes are not needed
    buf := make([]byte, 512)
    _, err := conn.Read(buf)
    if err != nil {
      // TODO factor in EOF errors and ignore those
      return errors.New(fmt.Sprintf("connection encountered error %s", err.Error()))
    }
    msg := deserialize(buf)

    fmt.Println(msg)
  }
}

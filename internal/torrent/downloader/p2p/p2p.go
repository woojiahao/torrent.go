package p2p

import (
  "errors"
  "fmt"
  "log"
  "net"
)

func Peer2Peer(conn net.Conn) error {
  for count := 0; ; count++ {
    // TODO Alter the byte length to be smaller if large bytes are not needed
    buf := make([]byte, 512)
    _, err := conn.Read(buf)
    if err != nil {
      // TODO factor in EOF errors and ignore those
      return errors.New(fmt.Sprintf("connection encountered error %s", err.Error()))
    }
    msg := Deserialize(buf)
    if msg.MessageID == BitfieldID {
      if count != 0 {
        log.Fatal("message received from cannot be bitfield as this is not the first message immediately after")
      }

      fmt.Println(msg.Payload)
    }
  }
}

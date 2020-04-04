package p2p

import (
  "errors"
  "fmt"
  "log"
  "net"
)

func Peer2Peer(conn net.Conn) error {
  var bitfield Bitfield
  isChoked, isInterested, isFirst := false, true, true

  fmt.Println(isChoked)

  for {
    // TODO Alter the byte length to be smaller if large bytes are not needed
    buf := make([]byte, 512)
    _, err := conn.Read(buf)
    if err != nil {
      // TODO factor in EOF errors and ignore those
      return errors.New(fmt.Sprintf("connection encountered error %s", err.Error()))
    }
    msg := Deserialize(buf)

    switch msg.MessageID {
    case ChokeID:
      // If the server is choked, then send a message to inform them that you are still interested
      if isInterested {
        interestedMessage := Message{MessageID: InterestedID}
        _, _ = conn.Write(interestedMessage.Serialize())
      } else {
        break
      }
    case BitfieldID:
      if !isFirst {
        log.Fatal("message received from cannot be bitfield as this is not the first message immediately after")
      }

      bitfield = msg.Payload
      fmt.Println(bitfield)
    }

    isFirst = false
  }
}

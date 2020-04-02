package p2p

import (
  "encoding/binary"
)

type message struct {
  lengthPrefix int
  messageID
  payload []byte
}

func (m *message) serialize() []byte {
  return []byte{}
}

func deserialize(b []byte) *message {
  lengthPrefix, messageID := int(binary.BigEndian.Uint32(b[:4])), messageID(int(b[4]))
  payloadSize, payload := 0, make([]byte, 0)
  switch messageID {
  case have:
    payloadSize = 4
  case bitfield:
    payloadSize = lengthPrefix - 1
  case piece:
    payloadSize = lengthPrefix - 9
  case port:
    payloadSize = 2
  case request:
  case cancel:
    payloadSize = 12
  }
  if payloadSize != 0 {
    payload = b[5 : 5+payloadSize]
  }
  return &message{
    lengthPrefix,
    messageID,
    payload,
  }
}

package p2p

import (
  "encoding/binary"
)

// For all integers in the payload, they will be regarded as BigEndian integers with 4 bytes
type message struct {
  lengthPrefix int
  messageID
  payload []byte
}

// This variable is used to send a keep alive packet to the server that cannot be serialized
var keepAlive = []byte{0, 0, 0, 0}

// Serializes a message into a stream of bytes. The given lengthPrefix is ignored as it must be calculated
// given the messageID and provided payload.
func (m *message) serialize() []byte {
  buf := make([]byte, 0)
  var length int
  switch m.messageID {
  case choke:
  case unchoke:
  case interested:
  case notInterested:
    length = 1
  case cancel:
  case request:
    length = 13
  case have:
    length = 5
  case bitfield:
  case piece:
    // For piece, the payload will be 8 bytes + block
    length = len(m.payload) + 1
  case port:
    length = 3
  }

  lengthPrefix := make([]byte, 4)
  binary.BigEndian.PutUint16(lengthPrefix, uint16(length))

  copy(buf, lengthPrefix)

  return buf
}

func deserialize(b []byte) *message {
  lengthPrefix, messageID := int(binary.BigEndian.Uint32(b[:4])), messageID(int(b[4]))
  payloadSize, payload := 0, make([]byte, 0)
  switch messageID {
  case have:
    payloadSize = 4
  case bitfield:
  case piece:
    payloadSize = lengthPrefix - 1
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

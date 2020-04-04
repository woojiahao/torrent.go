package p2p

import (
  . "github.com/woojiahao/torrent.go/internal/utility"
)

// Piece and Bitfield do not have a default length prefixes and payload sizes
// because they have dynamic sizes that have to be calculated separately
var (
  lengthPrefixes = map[MessageID]int{
    ChokeID:         1,
    UnchokeID:       1,
    InterestedID:    1,
    NotInterestedID: 1,
    CancelID:        13,
    RequestID:       13,
    HaveID:          5,
    PortID:          3,
  }

  payloadSizes = map[MessageID]int{
    ChokeID:         0,
    UnchokeID:       0,
    InterestedID:    0,
    NotInterestedID: 0,
    CancelID:        12,
    RequestID:       12,
    HaveID:          4,
    PortID:          2,
  }

  // This variable is used to send a keep alive packet to the server that cannot be serialized
  KeepAlive = []byte{0, 0, 0, 0}
)

// For all integers in the payload, they will be regarded as BigEndian integers with 4 bytes
type Message struct {
  LengthPrefix int
  MessageID    MessageID
  Payload      []byte
}

// Serializes a message into a stream of bytes. The given lengthPrefix is ignored as it must be calculated
// given the MessageID and provided payload.
func (m *Message) Serialize() []byte {
  messageID := m.MessageID
  buf := make([]byte, 0)

  var length int
  if messageID == PieceID || messageID == BitfieldID {
    length = len(m.Payload) + 1
  } else {
    length = lengthPrefixes[messageID]
  }

  lengthPrefix := ToBigEndian(length, 4)

  buf = append(buf, lengthPrefix...)
  buf = append(buf, byte(int(m.MessageID)))
  if m.Payload != nil {
    buf = append(buf, m.Payload...)
  }

  return buf
}

func Deserialize(b []byte) *Message {
  lengthPrefix, messageID := FromBigEndian(b[:4]), MessageID(int(b[4]))

  var payloadSize int
  if messageID == PieceID || messageID == BitfieldID {
    payloadSize = lengthPrefix - 1
  } else {
    payloadSize = payloadSizes[messageID]
  }

  payload := make([]byte, 0)
  if payloadSize != 0 {
    payload = b[5 : 5+payloadSize]
  }

  return &Message{
    lengthPrefix,
    messageID,
    payload,
  }
}

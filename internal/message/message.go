package message

import (
  "fmt"
  "github.com/woojiahao/torrent.go/internal/connection"
  . "github.com/woojiahao/torrent.go/internal/utility"
  "io"
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

func New(id MessageID, payload ...byte) *Message {
  return &Message{MessageID: id, Payload: payload}
}

func Deserialize(b []byte) *Message {
  lengthPrefix := FromBigEndian(b[:4])

  // If keep alive
  if lengthPrefix == 0 {
    return nil
  }

  messageID := MessageID(int(b[4]))

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

func Read(c *connection.Connection) (*Message, error) {
  lengthBuf := make([]byte, 4)
  _, err := io.ReadFull(c.Conn, lengthBuf)
  if err != nil {
    return nil, err
  }
  length := FromBigEndian(lengthBuf)

  if length == 0 {
    return nil, nil
  }

  buf, err := c.Receive(length)
  if err != nil {
    return nil, err
  }

  fullMessage := make([]byte, length+4)
  copy(fullMessage[:4], lengthBuf)
  copy(fullMessage[4:], buf)

  return Deserialize(fullMessage), nil
}

// Reads a PIECE message payload into the buffer
// Copies the payload to the buffer
func (m *Message) ParseBlock(index int, buf []byte) (int, error) {
  if m.MessageID != PieceID {
    return 0, fmt.Errorf("message is not of type PIECE")
  }

  payload := m.Payload

  if len(payload) < 8 {
    return 0, fmt.Errorf("payload is an invalid length")
  }

  blockIndex, begin, block := FromBigEndian(payload[:4]), FromBigEndian(payload[4:8]), payload[8:]

  if blockIndex != index {
    return 0, fmt.Errorf("expected index %d, received %d instead", index, blockIndex)
  }

  // The begin either points to the end of the buffer or further
  if begin >= len(buf) {
    return 0, fmt.Errorf("block offset cannot be greater than or equal to the piece buffer size")
  }

  // If the offset + payload exceeds the piece buffer's size
  if begin+len(block) > len(buf) {
    return 0, fmt.Errorf("payload too large for offset %d with piece buffer length of %d", begin, len(buf))
  }

  copy(buf[begin:], buf)
  return len(block), nil
}

func (m *Message) ParseHave() (int, error) {
  if m.MessageID != HaveID {
    return 0, fmt.Errorf("message is not of type HAVE")
  }

  if len(m.Payload) > 4 {
    return 0, fmt.Errorf("HAVE messages should have payloads of length 4")
  }

  index := FromBigEndian(m.Payload)

  return index, nil
}

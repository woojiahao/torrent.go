package message

import (
  "github.com/woojiahao/torrent.go/internal/message"
  . "github.com/woojiahao/torrent.go/internal/utility"
  "testing"
)

// Format: <length>:<id>:<payload>

// Choke 1:0
func TestChokeDeserialization(t *testing.T) {
  testDeserialization(t, 1, message.ChokeID)
}

// Unchoke 1:1
func TestUnchokeDeserialization(t *testing.T) {
  testDeserialization(t, 1, message.UnchokeID)
}

// Interested 1:2
func TestInterestedDeserialization(t *testing.T) {
  testDeserialization(t, 1, message.InterestedID)
}

// NotInterested 1:3
func TestNotInterestedDeserialization(t *testing.T) {
  testDeserialization(t, 1, message.NotInterestedID)
}

// Have 5:4:<piece index>
func TestHaveDeserialization(t *testing.T) {
  testHave(t, testDeserialization)
}

// Bitfield 1+X:5:<bitfield>
func TestBitfieldDeserialization(t *testing.T) {
  testBitfield(t, testDeserialization)
}

// Request 13:6:<index><begin><length>
func TestRequestDeserialization(t *testing.T) {
  testWithPayload(t, 13, message.RequestID, testDeserialization)
}

// Piece 9+X:7:<index><begin><block>
func TestPieceDeserialization(t *testing.T) {
  testPiece(t, testDeserialization)
}

// Cancel 13:8:<index><begin><length>
func TestCancelDeserialization(t *testing.T) {
  testWithPayload(t, 13, message.CancelID, testDeserialization)
}

// Port 3:9:<listen port>
func TestPortDeserialization(t *testing.T) {
  testPiece(t, testDeserialization)
}

func buildMessage(lengthPrefix int, id message.MessageID, payload []byte) *message.Message {
  return &message.Message{lengthPrefix, id, payload}
}

func buildMessageBytes(length int, id message.MessageID, payload []byte) []byte {
  buf := make([]byte, 0)
  lengthPrefix := ToBigEndian(length, 4)
  buf = append(buf, lengthPrefix...)
  buf = append(buf, byte(int(id)))
  buf = append(buf, payload...)
  return buf
}

func testDeserialization(t *testing.T, lengthPrefix int, id message.MessageID, payload ...byte) {
  expected := buildMessage(lengthPrefix, id, payload)
  mb := buildMessageBytes(lengthPrefix, id, payload)
  assertDeserialization(t, expected, message.Deserialize(mb))
}

func assertDeserialization(t *testing.T, expected, actual *message.Message) {
  if actual.LengthPrefix != expected.LengthPrefix {
    t.Errorf(
      "deserialization failed - mismatched length prefix; expected %v, got %v",
      expected.LengthPrefix,
      actual.LengthPrefix,
    )
  }

  if actual.MessageID != expected.MessageID {
    t.Errorf(
      "deserialization failed - mismatched message ID; expected %v, got %v",
      expected.MessageID,
      actual.MessageID,
    )
  }

  exPayload, acPayload := expected.Payload, actual.Payload

  if len(exPayload) != len(acPayload) {
    t.Errorf(
      "deserialization failed - payload length; expected %v, got %v",
      len(exPayload),
      len(acPayload),
    )
  }

  for i, b := range exPayload {
    if acPayload[i] != b {
      t.Errorf(
        "deserialization failed - byte mismatched; expected %v, got %v",
        b,
        acPayload[i],
      )
    }
  }
}

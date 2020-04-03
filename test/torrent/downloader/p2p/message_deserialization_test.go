package p2p

import (
  . "github.com/woojiahao/torrent.go/internal/torrent/downloader/p2p"
  . "github.com/woojiahao/torrent.go/internal/utility"
  "testing"
)

// Format: <length>:<id>:<payload>

// Choke 1:0
func TestChokeDeserialization(t *testing.T) {
  testDeserialization(t, 1, Choke)
}

// Unchoke 1:1
func TestUnchokeDeserialization(t *testing.T) {
  testDeserialization(t, 1, Unchoke)
}

// Interested 1:2
func TestInterestedDeserialization(t *testing.T) {
  testDeserialization(t, 1, Interested)
}

// NotInterested 1:3
func TestNotInterestedDeserialization(t *testing.T) {
  testDeserialization(t, 1, NotInterested)
}

// Have 5:4:<piece index>
func TestHaveDeserialization(t *testing.T) {
  for index := 0; index < 9999; index++ {
    testDeserialization(t, 5, Have, ToBigEndian(index, 4)...)
  }
}

// Bitfield 1+X:5:<bitfield>
func TestBitfieldDeserialization(t *testing.T) {
  generateBitfield := func(initial byte, size int) []byte {
    bitfield := make([]byte, size)
    for i := range bitfield {
      bitfield[i] = initial
    }
    return bitfield
  }

  for size := 1; size < 50; size++ {
    for initial := 1; initial < 100; initial++ {
      bitfield := generateBitfield(byte(initial), size)
      testDeserialization(t, 1+len(bitfield), Bitfield, bitfield...)
    }
  }
}

// Request 13:6:<index><begin><length>
func TestRequestDeserialization(t *testing.T) {
  testDeserializationWithPayload(t, 13, Request)
}

// Piece 9+X:7:<index><begin><block>
func TestPieceDeserialization(t *testing.T) {
  block := []byte{25, 69, 112, 187, 115, 1, 0, 255, 199, 100, 1, 0}

  generatePiece := func(index, begin int, block []byte) []byte {
    buf := make([]byte, 0, 8+len(block))
    buf = append(buf, ToBigEndian(index, 4)...)
    buf = append(buf, ToBigEndian(begin, 4)...)
    buf = append(buf, block...)
    return buf
  }

  for index := 1; index < 99; index++ {
    for begin := 1; begin < 99; begin++ {
      piece := generatePiece(index, begin, block)
      testDeserialization(t, 9+len(block), Piece, piece...)
    }
  }
}

// Cancel 13:8:<index><begin><length>
func TestCancelDeserialization(t *testing.T) {
  testDeserializationWithPayload(t, 13, Cancel)
}

// Port 3:9:<listen port>
func TestPortDeserialization(t *testing.T) {
  for port1 := 0; port1 < 100; port1++ {
    for port2 := 0; port2 < 100; port2++ {
      testDeserialization(t, 3, Port, byte(port1), byte(port2))
    }
  }
}

func buildMessage(lengthPrefix int, id MessageID, payload []byte) *Message {
  return &Message{lengthPrefix, id, payload}
}

func buildMessageBytes(length int, id MessageID, payload []byte) []byte {
  buf := make([]byte, 0)
  lengthPrefix := ToBigEndian(length, 4)
  buf = append(buf, lengthPrefix...)
  buf = append(buf, byte(int(id)))
  buf = append(buf, payload...)
  return buf
}

func testDeserialization(t *testing.T, lengthPrefix int, id MessageID, payload ...byte) {
  expected := buildMessage(lengthPrefix, id, payload)
  mb := buildMessageBytes(lengthPrefix, id, payload)
  assertDeserialization(t, expected, Deserialize(mb))
}

// Test deserialization of message IDs with payload of <index><begin><length>
func testDeserializationWithPayload(t *testing.T, lengthPrefix int, id MessageID) {
  generatePayload := func(index, begin, length int) []byte {
    buf := make([]byte, 0, 12)
    buf = append(buf, ToBigEndian(index, 4)...)
    buf = append(buf, ToBigEndian(begin, 4)...)
    buf = append(buf, ToBigEndian(length, 4)...)
    return buf
  }

  for index := 1; index <= 100; index++ {
    for begin := 1; begin <= 100; begin++ {
      for length := 1; length <= 100; length++ {
        payload := generatePayload(index, begin, length)
        testDeserialization(t, lengthPrefix, id, payload...)
      }
    }
  }
}

func assertDeserialization(t *testing.T, expected, actual *Message) {
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

package p2p

import (
  "github.com/woojiahao/torrent.go/internal/message"
  . "github.com/woojiahao/torrent.go/internal/utility"
  "testing"
)

func TestChokeSerialization(t *testing.T) {
  testSerialization(t, 1, message.ChokeID)
}

func TestUnchokeSerialization(t *testing.T) {
  testSerialization(t, 1, message.UnchokeID)
}

func TestInterestedSerialization(t *testing.T) {
  testSerialization(t, 1, message.InterestedID)
}

func TestNotInterestedSerialization(t *testing.T) {
  testSerialization(t, 1, message.NotInterestedID)
}

func TestHaveSerialization(t *testing.T) {
  testSerialization(t, 1, message.NotInterestedID)
}

func TestBitfieldSerialization(t *testing.T) {
  testBitfield(t, testSerialization)
}

func TestRequestSerialization(t *testing.T) {
  testWithPayload(t, 13, message.RequestID, testSerialization)
}

func TestPieceSerialization(t *testing.T) {
  testPiece(t, testSerialization)
}

func TestCancelSerialization(t *testing.T) {
  testWithPayload(t, 13, message.CancelID, testSerialization)
}

func TestPortSerialization(t *testing.T) {
  testPort(t, testSerialization)
}

func testSerialization(t *testing.T, length int, id message.MessageID, payload ...byte) {
  expected := make([]byte, 0)
  expected = append(expected, ToBigEndian(length, 4)...)
  expected = append(expected, byte(int(id)))
  expected = append(expected, payload...)

  m := message.Message{
    LengthPrefix: length,
    MessageID:    id,
    Payload:      payload,
  }

  assertSerialization(t, expected, m.Serialize())
}

func assertSerialization(t *testing.T, expected, actual []byte) {
  if len(expected) != len(actual) {
    t.Errorf("message length not equal; expected: %d, got: %d", len(expected), len(actual))
  }

  for i, d := range expected {
    if d != actual[i] {
      t.Errorf("byte at index %d does not match expected; expected: %v, got: %v", i, d, actual[i])
    }
  }
}

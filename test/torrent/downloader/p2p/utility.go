package p2p

import (
  . "github.com/woojiahao/torrent.go/internal/downloader/p2p"
  . "github.com/woojiahao/torrent.go/internal/utility"
  "testing"
)

type testType func(*testing.T, int, MessageID, ...byte)

var block = []byte{25, 69, 112, 187, 115, 1, 0, 255, 199, 100, 1, 0}

// Test with payload of <index><begin><length>
func testWithPayload(t *testing.T, lengthPrefix int, id MessageID, test testType) {
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
        test(t, lengthPrefix, id, payload...)
      }
    }
  }
}

func testHave(t *testing.T, test testType) {
  for index := 0; index < 9999; index++ {
    test(t, 5, HaveID, ToBigEndian(index, 4)...)
  }
}

func testPiece(t *testing.T, test testType) {
  generatePiece := func(index, begin int) []byte {
    buf := make([]byte, 0, 8+len(block))
    buf = append(buf, ToBigEndian(index, 4)...)
    buf = append(buf, ToBigEndian(begin, 4)...)
    buf = append(buf, block...)
    return buf
  }

  for index := 1; index < 99; index++ {
    for begin := 1; begin < 99; begin++ {
      piece := generatePiece(index, begin)
      test(t, 9+len(block), PieceID, piece...)
    }
  }
}

func testBitfield(t *testing.T, test testType) {
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
      test(t, 1+len(bitfield), BitfieldID, bitfield...)
    }
  }
}

func testPort(t *testing.T, test testType) {
  for port1 := 0; port1 < 100; port1++ {
    for port2 := 0; port2 < 100; port2++ {
      test(t, 3, PortID, byte(port1), byte(port2))
    }
  }
}

package p2p

import (
  . "github.com/woojiahao/torrent.go/internal/torrent/downloader/p2p"
  . "github.com/woojiahao/torrent.go/internal/utility"
  "testing"
)

func TestBitfieldHasPiece(t *testing.T) {
  for i := 1; i < 999; i++ {
    bitfield := generateBitfield(i)
    first := bitfield.HasPiece(0)
    middle := bitfield.HasPiece(i / 2)
    oneAfterLast := bitfield.HasPiece(i+1)

    if !first {
      t.Errorf("first should always be true")
    }

    if !middle {
      t.Errorf("middle should always be true")
    }

    if oneAfterLast {
      t.Errorf("one after last cannot be true")
    }
  }
}

// Generates a bitfield based on the number of bits to be filled
func generateBitfield(filledBits int) Bitfield {
  generateByte := func(bits int) byte {
    b := 0
    for i := 7; i >= 8-bits; i-- {
      b += Pow(2, i)
    }
    return byte(b)
  }

  chunkSize, bitfield := filledBits/8, make([]byte, 0)

  for i, remainingBits := 0, filledBits; i < chunkSize || remainingBits > 0; i++ {
    use := Min(remainingBits, 8)
    bitfield = append(bitfield, generateByte(use))
    remainingBits -= use
  }

  return bitfield
}

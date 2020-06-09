package bitfield

import (
  "github.com/woojiahao/torrent.go/internal/bitfield"
  . "github.com/woojiahao/torrent.go/internal/utility"
  "testing"
)

func TestBitfieldHasPiece(t *testing.T) {
  for i := 1; i < 999; i++ {
    bf := generateBitfield(i)
    first := bf.HasPiece(0)
    middle := bf.HasPiece(i / 2)
    oneAfterLast := bf.HasPiece(i+1)

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
func generateBitfield(filledBits int) bitfield.Bitfield {
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

package p2p

import (
  . "github.com/woojiahao/torrent.go/internal/utility"
)

// A bitfield is an array of bytes with each byte
// Each bit in the bitfield corresponds to a piece index
type Bitfield []byte

// pieceIndex starts with 0
func (b Bitfield) HasPiece(pieceIndex int) bool {
  // Which byte the piece resides in
  byteIndex := pieceIndex / 8
  offset := pieceIndex % 8
  mask := byte(Pow(2, 7-offset))
  if len(b)-1 < byteIndex {
    return false
  }
  return b[byteIndex]&mask == mask
}

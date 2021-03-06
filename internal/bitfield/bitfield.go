package bitfield

import "fmt"

// A bitfield is an array of bytes with each byte
// Each bit in the bitfield corresponds to a piece index
type Bitfield []byte

// pieceIndex starts with 0
// We first shift the byte to the right such that the piece will be at the
// lowest bit. Then we simply mask the byte with 1 and if the result is 1,
// then we know that the piece exists, else, the piece does not exist
func (b Bitfield) HasPiece(pieceIndex int) bool {
  byteIndex := pieceIndex / 8
  offset := pieceIndex % 8
  if byteIndex < 0 || len(b)-1 < byteIndex {
    return false
  }
  return b[byteIndex]>>(7-offset)&1 != 0
}

// Sets a specified piece index to 1
func (b Bitfield) SetPiece(pieceIndex int) {
  fmt.Printf("piece index is %d\n", pieceIndex)
  byteIndex := pieceIndex / 8
  fmt.Printf("byte index is %d\n", byteIndex)
  offset := pieceIndex % 8
  fmt.Printf("offset is %d\n", offset)
  if byteIndex < 0 || len(b)-1 < byteIndex {
    return
  }
  b[byteIndex] |= 1 << (7 - offset)
}

package piece

import (
  "errors"
  "fmt"
  . "github.com/woojiahao/torrent.go/internal/utility"
)

type Pieces [][20]byte

const pieceSize = 20

// Generate the pieces of a torrent file
func CreatePieces(piecesStr string) Pieces {
  piecesCount := len(piecesStr) / pieceSize
  pieces := make([][pieceSize]byte, piecesCount)

  if len(piecesStr)%pieceSize != 0 {
    Check(errors.New(fmt.Sprintf("invalid pieces format; not a multiple of %d", pieceSize)))
  }

  for i := 0; i < piecesCount; i += pieceSize {
    byteSlice := []byte(piecesStr[i : i+pieceSize])
    var byteChunk [pieceSize]byte
    copy(byteChunk[:], byteSlice[:pieceSize])
    pieces[i] = byteChunk
  }

  return pieces
}

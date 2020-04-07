package torrent

import (
  "errors"
  "fmt"
  . "github.com/woojiahao/torrent.go/internal/utility"
)

type Pieces [][20]byte

// Generate the pieces of a torrent file
func createPieces(piecesStr string) Pieces {
  const pieceSize = 20
  piecesCount := len(piecesStr)/pieceSize
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

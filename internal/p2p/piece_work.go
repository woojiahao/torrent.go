package p2p

import (
  "bytes"
  "fmt"
  . "github.com/woojiahao/torrent.go/internal/utility"
)

// For a piece to be downloaded, we need to know its position in the file,
// its hash, and the length of this piece
type pieceWork struct {
  index  int
  hash   [20]byte
  length int
}

// Verifies that a downloaded piece matches the advertised piece hash
func (pw *pieceWork) checkIntegrity(piece []byte) error {
  pieceHash := GenerateSHA1Hash(string(piece)).Sum(nil)
  if bytes.Equal(pw.hash[:], pieceHash) {
    return fmt.Errorf("SHA1 hash of downloaded piece is not the same as the original piece SHA1 hash")
  }
  return nil
}

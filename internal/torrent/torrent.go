package torrent

type (
  Torrent interface {
    GetAnnounce() string
    GetLength() int
    GetPieces() Pieces
    GetPieceLength() int
  }
)

const (
  announce    = "announce"
  info        = "info"
  pieceLength = "piece length"
  name        = "name"
  pieces      = "pieces"
  length      = "length"
  files       = "files"
)

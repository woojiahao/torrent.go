package torrent

type (
  Pieces  [][20]byte
  Torrent interface {
    GetAnnounce() string
    GetLength() int
    GetPieces() Pieces
    GetPieceLength() int
  }
)

package torrent

type Torrent interface {
  GetAnnounce() string
  GetLength() int
  GetPieces() Pieces
  GetPieceLength() int
}

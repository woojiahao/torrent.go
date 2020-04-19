package torrent

import "github.com/woojiahao/torrent.go/internal/piece"

type (
  TorrentFile interface {
    GetName() string
    GetAnnounce() string
    GetLength() int
    GetPieces() piece.Pieces
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

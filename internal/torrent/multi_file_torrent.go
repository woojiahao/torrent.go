package torrent

import "strings"

// Multi file torrent structures
type (
  multiFileTorrent struct {
    announce string
    info     multiFileInfo
  }

  multiFileInfo struct {
    files       []file
    name        string
    pieceLength int
    Pieces
  }

  file struct {
    length int
    paths  []string
  }
)

func (t multiFileTorrent) GetAnnounce() string {
  return t.announce
}

func (t multiFileTorrent) GetLength() int {
  return 0
}

func (t multiFileTorrent) GetPieces() Pieces {
  return t.info.Pieces
}

func (t multiFileTorrent) GetPieceLength() int {
  return t.info.pieceLength
}

func (f file) path() string {
  return strings.Join(f.paths, "/")
}

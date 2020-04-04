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
    pieces
  }

  file struct {
    length int
    paths  []string
  }
)

func (t multiFileTorrent) getAnnounce() string {
  return t.announce
}

func (t multiFileTorrent) getLength() int {
  return 0
}

func (t multiFileTorrent) getPieces() pieces {
  return t.info.pieces
}

func (t multiFileTorrent) getPieceLength() int {
  return t.info.pieceLength
}

func (f file) path() string {
  return strings.Join(f.paths, "/")
}

package torrent

// Single file torrent structures
type (
  singleFileTorrent struct {
    announce string
    info     singleFileInfo
  }

  singleFileInfo struct {
    length      int
    name        string
    pieceLength int
    Pieces
  }
)

func (t singleFileTorrent) GetAnnounce() string {
  return t.announce
}

func (t singleFileTorrent) GetLength() int {
  return t.info.length
}

func (t singleFileTorrent) GetPieces() Pieces {
  return t.info.Pieces
}

func (t singleFileTorrent) GetPieceLength() int {
  return t.info.pieceLength
}

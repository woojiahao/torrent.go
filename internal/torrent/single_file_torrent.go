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
    pieces
  }
)

func (t singleFileTorrent) getAnnounce() string {
  return t.announce
}

func (t singleFileTorrent) getLength() int {
  return t.info.length
}

func (t singleFileTorrent) getPieces() pieces {
  return t.info.pieces
}

func (t singleFileTorrent) getPieceLength() int {
  return t.info.pieceLength
}

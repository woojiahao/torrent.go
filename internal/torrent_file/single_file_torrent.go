package torrent_file

// Single file torrent structures
type (
  singleFileTorrentFile struct {
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

func (t singleFileTorrentFile) GetAnnounce() string {
  return t.announce
}

func (t singleFileTorrentFile) GetLength() int {
  return t.info.length
}

func (t singleFileTorrentFile) GetPieces() Pieces {
  return t.info.Pieces
}

func (t singleFileTorrentFile) GetPieceLength() int {
  return t.info.pieceLength
}

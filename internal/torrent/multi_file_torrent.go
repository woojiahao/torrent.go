package torrent

import (
  . "github.com/woojiahao/torrent.go/internal/bencoding"
  "strings"
)

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

// Create the files list for multi file torrents
func parseFiles(filesLst TList) []file {
  files := make([]file, 0)

  for _, f := range filesLst {
    data := ToDict(f)

    pathsLst, paths := ToList(data["path"]), make([]string, 0)
    for _, p := range pathsLst {
      paths = append(paths, ToString(p).Value())
    }

    file := file{
      length: ToInt(data["length"]).Value(),
      paths:  paths,
    }
    files = append(files, file)
  }

  return files
}

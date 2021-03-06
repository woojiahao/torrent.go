package torrent

import (
  . "github.com/woojiahao/torrent.go/internal/bencoding"
  . "github.com/woojiahao/torrent.go/internal/piece"
  "strings"
)

// Multi file torrent structures
type (
  multiFileTorrentFile struct {
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

func (t multiFileTorrentFile) GetName() string {
  return t.info.name
}

func (t multiFileTorrentFile) GetAnnounce() string {
  return t.announce
}

func (t multiFileTorrentFile) GetLength() int {
  length := 0
  for _, file := range t.info.files {
    length += file.length
  }
  return length
}

func (t multiFileTorrentFile) GetPieces() Pieces {
  return t.info.Pieces
}

func (t multiFileTorrentFile) GetPieceLength() int {
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

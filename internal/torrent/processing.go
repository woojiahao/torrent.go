package torrent

import (
  "errors"
  "fmt"
  . "github.com/woojiahao/torrent.go/internal/bencoding"
  . "github.com/woojiahao/torrent.go/internal/utility"
)


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

// Parses a torrent file into either a single file torrent or multi file torrent
func parseTorrentFile(torrentMetadata TDict) (Torrent, bool) {
  announce, info := ToString(torrentMetadata["announce"]).Value(),
    ToDict(torrentMetadata["info"])

  isSingle := info["files"] == nil

  name, pieceLength, pieces := ToString(info["name"]).Value(),
    ToInt(info["piece length"]).Value(),
    createPieces(ToString(info["pieces"]).Value())

  var torrent Torrent

  if isSingle {
    torrent = singleFileTorrent{
      announce,
      singleFileInfo{
        length:      ToInt(info["length"]).Value(),
        name:        name,
        pieceLength: pieceLength,
        Pieces:      pieces,
      },
    }
  } else {
    torrent = multiFileTorrent{
      announce,
      multiFileInfo{
        files:       parseFiles(ToList(info["files"])),
        name:        name,
        pieceLength: pieceLength,
        Pieces:      pieces,
      },
    }
  }
  return torrent, isSingle
}

package torrent

import (
  . "github.com/woojiahao/torrent.go/internal/bencoding"
  . "github.com/woojiahao/torrent.go/internal/utility"
  "strings"
)

type (
  pieces  [][20]byte
  torrent interface{}
)

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

func (f file) path() string {
  return strings.Join(f.paths, "/")
}

// Generate the pieces of a torrent file
func createPieces(piecesStr string) pieces {
  pieces := make([][20]byte, 0)

  for i := 0; i < len(piecesStr); i += 20 {
    byteSlice := []byte(piecesStr[i : i+20])
    var byteChunk [20]byte
    copy(byteChunk[:], byteSlice[:20])
    pieces = append(pieces, byteChunk)
  }

  return pieces
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

// Parses a torrent file into either a single file torrent or multi file torrent
func parseTorrentFile(torrentFileContent TDict) (torrent, bool) {
  announce, info := ToString(torrentFileContent["announce"]).Value(),
    ToDict(torrentFileContent["info"])

  isSingle := info["files"] == nil

  name, pieceLength, pieces := ToString(info["name"]).Value(),
    ToInt(info["piece length"]).Value(),
    createPieces(ToString(info["pieces"]).Value())

  var torrent torrent

  if isSingle {
    torrent = singleFileTorrent{
      announce,
      singleFileInfo{
        length:      ToInt(info["length"]).Value(),
        name:        name,
        pieceLength: pieceLength,
        pieces:      pieces,
      },
    }
  } else {
    torrent = multiFileTorrent{
      announce,
      multiFileInfo{
        files:       parseFiles(ToList(info["files"])),
        name:        name,
        pieceLength: pieceLength,
        pieces:      pieces,
      },
    }
  }
  return torrent, isSingle
}

// Downloads a torrent from the given file path
func Download(torrentFilename string) {

  if NotExist(torrentFilename) {
    panic("file does not exist")
  } else if IsDir(torrentFilename) {
    panic("filename points to a directory")
  } else if !IsFileType(torrentFilename, "torrent") {
    panic("given filename must be a torrent file")
  }

  fileContents := ReadFileContents(torrentFilename)

  torrentMetadata := ToDict(Decode(fileContents))

  torrent, isSingle := parseTorrentFile(torrentMetadata)

  info := torrentMetadata["info"].Encode()

  if isSingle {
    torrent, ok := torrent.(singleFileTorrent)
    if !ok {
      panic("failed to convert to single file torrent")
    }

    RequestTracker(torrent.announce, info, torrent.info.length)
  } else {
    torrent, ok := torrent.(multiFileTorrent)
    if !ok {
      panic("failed to convert to multi-file torrent")
    }

    RequestTracker(torrent.announce, info, 0)
  }
}

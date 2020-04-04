package torrent

import (
  "errors"
  "fmt"
  . "github.com/woojiahao/torrent.go/internal/bencoding"
  "github.com/woojiahao/torrent.go/internal/torrent/downloader"
  "github.com/woojiahao/torrent.go/internal/torrent/tracker"
  . "github.com/woojiahao/torrent.go/internal/utility"
  "log"
)

type (
  pieces  [][20]byte
  torrent interface {
    getAnnounce() string
    getLength() int
    getPieces() pieces
    getPieceLength() int
  }
)

// Generate the pieces of a torrent file
func createPieces(piecesStr string) pieces {
  const pieceSize = 20
  pieces := make([][pieceSize]byte, 0)

  if len(piecesStr)%pieceSize != 0 {
    LogCheck(errors.New(fmt.Sprintf("invalid pieces format; not a multiple of %d", pieceSize)))
  }

  for i := 0; i < len(piecesStr)/pieceSize; i += pieceSize {
    byteSlice := []byte(piecesStr[i : i+pieceSize])
    var byteChunk [pieceSize]byte
    copy(byteChunk[:], byteSlice[:pieceSize])
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
func parseTorrentFile(torrentMetadata TDict) (torrent, bool) {
  announce, info := ToString(torrentMetadata["announce"]).Value(),
    ToDict(torrentMetadata["info"])

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

// TODO Retry the download if the original returns a failure
// Downloads a torrent from the given file path
func Download(torrentFilename string) {

  log.Printf("starting torrent download")

  if NotExist(torrentFilename) {
    LogCheck(&fileError{torrentFilename, "does not exist"})
  } else if IsDir(torrentFilename) {
    LogCheck(&fileError{torrentFilename, "points to a directory"})
  } else if !IsFileType(torrentFilename, "torrent") {
    LogCheck(&fileError{torrentFilename, "is not a .torrent file"})
  }

  log.Print("downloading torrent file contents")
  fileContents := ReadFileContents(torrentFilename)

  log.Print("decoding torrent metadata")
  torrentMetadata := ToDict(Decode(fileContents))

  log.Print("parsing torrent metadata into torrent file")
  torrent, _ := parseTorrentFile(torrentMetadata)

  log.Printf("pieces available %d with piece length of %d", len(torrent.getPieces()), torrent.getPieceLength())

  info := torrentMetadata["info"].Encode()

  log.Print("requesting tracker for information")
  peers, infoHash, peerID := tracker.RequestTracker(torrent.getAnnounce(), info, torrent.getLength())

  log.Print("downloading torrent with tracker information")
  downloader.Download(peers, infoHash, peerID)
}

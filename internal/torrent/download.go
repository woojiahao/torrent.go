package torrent

import (
  . "github.com/woojiahao/torrent.go/internal/bencoding"
  "github.com/woojiahao/torrent.go/internal/downloader"
  "github.com/woojiahao/torrent.go/internal/tracker"
  . "github.com/woojiahao/torrent.go/internal/utility"
  "log"
)

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

// TODO Retry the download if the original returns a failure
// Downloads a torrent from the given file path
func Download(torrentFilename string) {

  log.Printf("starting torrent download")

  if NotExist(torrentFilename) {
    Check(&fileError{torrentFilename, "does not exist"})
  } else if IsDir(torrentFilename) {
    Check(&fileError{torrentFilename, "points to a directory"})
  } else if !IsFileType(torrentFilename, "torrent") {
    Check(&fileError{torrentFilename, "is not a .torrent file"})
  }

  log.Print("downloading torrent file contents")
  fileContents := ReadFileContents(torrentFilename)

  log.Print("decoding torrent metadata")
  torrentMetadata := ToDict(Decode(fileContents))

  log.Print("parsing torrent metadata into torrent file")
  torrent, _ := parseTorrentFile(torrentMetadata)

  log.Print("requesting tracker for information")
  peers, infoHash, peerID := tracker.RequestTracker(
    torrent.GetAnnounce(),
    torrentMetadata["info"].Encode(),
    torrent.GetLength(),
  )

  log.Print("downloading torrent with tracker information")
  downloader.Download(peers, infoHash, peerID)
}

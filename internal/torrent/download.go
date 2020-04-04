package torrent

import (
  . "github.com/woojiahao/torrent.go/internal/bencoding"
  "github.com/woojiahao/torrent.go/internal/downloader"
  "github.com/woojiahao/torrent.go/internal/tracker"
  . "github.com/woojiahao/torrent.go/internal/utility"
  "log"
)

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

  log.Print("requesting tracker for information")
  peers, infoHash, peerID := tracker.RequestTracker(
    torrent.GetAnnounce(),
    torrentMetadata["info"].Encode(),
    torrent.GetLength(),
  )

  log.Print("downloading torrent with tracker information")
  downloader.Download(peers, torrent, infoHash, peerID)
}

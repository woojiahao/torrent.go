package torrent

import (
  . "github.com/woojiahao/torrent.go/internal/bencoding"
  "github.com/woojiahao/torrent.go/internal/downloader"
  "github.com/woojiahao/torrent.go/internal/tracker"
  . "github.com/woojiahao/torrent.go/internal/utility"
  "io/ioutil"
  "os"
  "path/filepath"
)

// Parses a torrent file into either a single file torrent or multi file torrent
func parseTorrentFile(torrentMetadata TDict) (Torrent, bool) {
  announce, info := ToString(torrentMetadata[announce]).Value(),
    ToDict(torrentMetadata[info])


  name, pieceLength, pieces := ToString(info[name]).Value(),
    ToInt(info[pieceLength]).Value(),
    createPieces(ToString(info[pieces]).Value())

  var torrent Torrent

  isSingle := info["files"] == nil

  if isSingle {
    torrent = singleFileTorrent{
      announce,
      singleFileInfo{
        length:      ToInt(info[length]).Value(),
        name:        name,
        pieceLength: pieceLength,
        Pieces:      pieces,
      },
    }
  } else {
    torrent = multiFileTorrent{
      announce,
      multiFileInfo{
        files:       parseFiles(ToList(info[files])),
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
  validateTorrentFilename(torrentFilename)

  fileContents := readFileContents(torrentFilename)

  torrentMetadata := ToDict(Decode(fileContents))

  torrent, _ := parseTorrentFile(torrentMetadata)

  peers, infoHash, peerID := tracker.RequestTracker(
    torrent.GetAnnounce(),
    torrentMetadata[info].Encode(),
    torrent.GetLength(),
  )

  downloader.Download(peers, infoHash, peerID)
}

func validateTorrentFilename(filename string) {
  file, err := os.Stat(filename)

  if os.IsNotExist(err) {
    err = &fileError{filename, "does not exist"}
  } else if file.IsDir() {
    err = &fileError{filename, "points to a directory"}
  } else if filepath.Ext(filename) != ".torrent" {
    err = &fileError{filename, "is not a .torrent file"}
  }

  Check(err)
}

func readFileContents(filename string) string {
  data, err := ioutil.ReadFile(filename)
  Check(err)
  return string(data)
}

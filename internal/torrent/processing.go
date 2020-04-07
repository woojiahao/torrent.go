package torrent

import (
  . "github.com/woojiahao/torrent.go/internal/bencoding"
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

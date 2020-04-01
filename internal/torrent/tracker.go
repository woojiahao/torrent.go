package torrent

import (
  "crypto/sha1"
  "encoding/binary"
  . "github.com/woojiahao/torrent.go/internal/bencoding"
  . "github.com/woojiahao/torrent.go/internal/utility"
  "io/ioutil"
  "net"
  "net/http"
  "strconv"
  "strings"
)

type (
  trackerResponse struct {
    failureReason  string
    warningMessage string
    interval       int
    minInterval    int
    trackerId      string
    complete       int
    incomplete     int
    peers          []peer
  }

  peer struct {
    ip   string
    port int
  }
)

// The peer_id is a 20 character string that is randomly generated at the start of each download
func generatePeerID() string {
  peerID := make([]string, 20)
  for i := 0; i < 20; i++ {
    isDigitOrLetter := RandomInt(0, 2)
    switch isDigitOrLetter {
    case 0:
      peerID[i] = strconv.Itoa(RandomInt(0, 10))
    case 1:
      peerID[i] = string(RandomChar())
    default:
      panic("invalid int")
    }
  }

  return strings.Join(peerID, "")
}

// The info_hash is the SHA1 hash representation of the bencoding info portion of the metadata
// The SHA1 hash generated is 40 characters long for human reading, it is in fact a hex string
// The tracker must receive the URL-encoded version of the hex string
func generateInfoHash(info string) string {
  h := sha1.New()
  h.Write([]byte(info))
  return string(h.Sum(nil))
}

func parsePeersBinary(peersBinary string) []peer {
  if len(peersBinary)%6 != 0 {
    panic("invalid peers string")
  }

  peers := make([]peer, 0)
  for i := 0; i < len(peersBinary)/6; i += 6 {
    ip, port := peersBinary[i:i+4], peersBinary[i+4:i+6]
    peer := peer{
      net.IP(ip).String(),
      int(binary.BigEndian.Uint16([]byte(port))),
    }
    peers = append(peers, peer)
  }
  return peers
}

// Convert the tracker bencoding response to trackerResponse struct
func parseTrackerResponse(metadata TDict) trackerResponse {
  return trackerResponse{
    ToString(metadata["failure reason"]).Value(),
    ToString(metadata["warning message"]).Value(),
    ToInt(metadata["interval"]).Value(),
    ToInt(metadata["min interval"]).Value(),
    ToString(metadata["tracker id"]).Value(),
    ToInt(metadata["complete"]).Value(),
    ToInt(metadata["incomplete"]).Value(),
    parsePeersBinary(ToString(metadata["peers"]).Value()),
  }
}

// TODO Add support for UDP connections
// Requests information from the given tracker
func requestTracker(trackerURL, info string, length int) trackerResponse {
  infoHash := generateInfoHash(info)
  var resp *http.Response
  defer func() {
    if resp != nil {
      _ = resp.Body.Close()
    }
  }()

  for port := 6881; port <= 6889; port++ {
    parameters := QueryParameters{
      "info_hash":  infoHash,
      "peer_id":    generatePeerID(),
      "port":       strconv.Itoa(port),
      "uploaded":   "0",
      "downloaded": "0",
      "left":       strconv.Itoa(length),
      "compact":    "1",
    }

    resp = GET(trackerURL, parameters)

    // TODO Might need to make this more dynamic for checking status code
    if resp.StatusCode == 200 {
      break
    }
  }

  if resp == nil {
    panic("cannot connect to tracker; unable to get response from announce url")
  }

  body, err := ioutil.ReadAll(resp.Body)
  Check(err)

  trackerResponseMetadata := ToDict(Decode(string(body)))
  trackerResponse := parseTrackerResponse(trackerResponseMetadata)
  return trackerResponse
}

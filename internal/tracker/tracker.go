package tracker

import (
  "encoding/binary"
  "errors"
  "fmt"
  . "github.com/woojiahao/torrent.go/internal/bencoding"
  . "github.com/woojiahao/torrent.go/internal/utility"
  "io/ioutil"
  "net"
  "net/http"
  "strconv"
  "strings"
  "time"
)

type Response struct {
  failureReason  string
  warningMessage string
  interval       int
  minInterval    int
  trackerId      string
  complete       int
  incomplete     int
  peers          []Peer
}

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
      Check(errors.New("invalid randomly generated integer"))
    }
  }

  return strings.Join(peerID, "")
}

func parsePeersBinary(peersBinary string) []Peer {
  const peerSize = 6
  if len(peersBinary)%peerSize != 0 {
    Check(errors.New(fmt.Sprintf("invalid peers string; length must be a multiple of %d", peerSize)))
  }

  peers := make([]Peer, 0)
  for i := 0; i < len(peersBinary)/peerSize; i += peerSize {
    ip, port := peersBinary[i:i+4], peersBinary[i+4:i+6]
    peer := Peer{
      net.IP(ip).String(),
      int(binary.BigEndian.Uint16([]byte(port))),
    }
    peers = append(peers, peer)
  }
  return peers
}

// Convert the tracker bencoding response to Response struct
func parseTrackerResponse(metadata TDict) *Response {
  return &Response{
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

func buildTrackerURLParameters(infoHash, peerID string, port, length int) *QueryParameters {
  return &QueryParameters{
    "info_hash":  infoHash,
    "peer_id":    peerID,
    "port":       strconv.Itoa(port),
    "uploaded":   "0",
    "downloaded": "0",
    "left":       strconv.Itoa(length),
    "compact":    "1",
  }
}

func queryTracker(trackerURL, infoHash, peerID string, length int) *http.Response {
  var resp *http.Response
  for port := 6881; port <= 6889; port++ {
    parameters := buildTrackerURLParameters(infoHash, peerID, port, length)
    resp = GET(trackerURL, parameters)

    // TODO Might need to make this more dynamic for checking status code
    if resp.StatusCode == 200 {
      break
    }
  }

  return resp
}

// TODO Add support for UDP connections
// Requests information from the given tracker
func RequestTracker(trackerURL, info string, length int) ([]Peer, string, string) {
  infoHash := string(GenerateSHA1Hash(info).Sum(nil))
  var peerID string
  var resp *http.Response
  defer func() {
    if resp != nil {
      _ = resp.Body.Close()
    }
  }()

  // When retrying to make a connection, we will pause the execution for 5 seconds
  // in case the servers don't respond to rapid successions of queries
  retry := 0
  for resp == nil && retry < 3 {
    if retry != 0 {
      time.Sleep(ToSeconds(5))
    }
    peerID = generatePeerID()
    resp = queryTracker(trackerURL, infoHash, peerID, length)
    retry++
  }

  if resp == nil {
    Check(errors.New("cannot connect to tracker; unable to get response from announce url"))
  }

  body, err := ioutil.ReadAll(resp.Body)
  Check(err)

  trackerResponseMetadata := ToDict(Decode(string(body)))

  trackerResponse := parseTrackerResponse(trackerResponseMetadata)

  if trackerResponse.failureReason != "" {
    Check(errors.New(fmt.Sprintf("tracker failed with reason %s", trackerResponse.failureReason)))
  } else if len(trackerResponse.peers) == 0 {
    Check(errors.New("no peers were provided by the tracker"))
  }

  return trackerResponse.peers, infoHash, peerID
}

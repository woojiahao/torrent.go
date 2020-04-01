package torrent

import (
  "crypto/sha1"
  "fmt"
  . "github.com/woojiahao/torrent.go/internal/utility"
  "io/ioutil"
  "strconv"
  "strings"
)

// TODO Method for generating base ten ASCII numbers for uploaded and downloaded
func generateBaseTenASCII(num int) int {
  return num*10 + 2
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

// Requests information from the given tracker
func RequestTracker(trackerURL, info string, length int) {
  infoHash := generateInfoHash(info)

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

    resp := GET(trackerURL, parameters)
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    Check(err)

    fmt.Println(string(body))
  }

  panic("cannot connect to tracker")
}

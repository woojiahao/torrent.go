package p2p

import (
  "fmt"
  "github.com/woojiahao/torrent.go/internal/client"
  "github.com/woojiahao/torrent.go/internal/message"
  "io"
  "log"
)

// The progress we make while downloading a single piece
// index -> index of the piece
// client -> connection to the peer
// buf -> downloaded blocks of the piece
// downloaded -> confirmed amount of pieces downloaded
// requested -> pieces that we have already requested for; we split this up as we may have requested for something
// but did not manage to download it. requested also acts as the offset in the block
// backlog -> ??
type pieceProgress struct {
  index      int
  client     *client.Client
  buf        []byte
  downloaded int
  requested  int
  backlog    int
}

func (p *pieceProgress) read() error {
  log.Printf("piece progress read :: piece is %d\n", p.index)
  msg, err := message.Read(p.client.Conn)
  if err != nil {
    if err == io.EOF {
      log.Printf("server responded with empty message, re-trying query")
      return nil
    }
    return err
  }

  // If keep alive
  if msg == nil {
    return nil
  }

  log.Printf("message id is %d\n", msg.MessageID)

  switch msg.MessageID {
  case message.ChokeID:
    p.client.Choked = true
  case message.UnchokeID:
    p.client.Choked = false
  case message.HaveID:
    index, err := msg.ParseHave()
    if err != nil {
      return err
    }
    p.client.Bitfield.SetPiece(index)
  case message.PieceID:
    n, err := msg.ParseBlock(p.index, p.buf)
    if err != nil {
      return err
    }
    p.downloaded += n
    p.backlog--
  case message.PortID:
    log.Printf("message received is a new port %v\n", msg.Payload)
  default:
    panic(fmt.Sprintf("should not be receiving message id of %d", msg.MessageID))
  }

  return nil
}

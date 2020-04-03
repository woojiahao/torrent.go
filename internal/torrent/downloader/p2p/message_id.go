package p2p

type MessageID int

const (
  Choke         MessageID = 0
  Unchoke       MessageID = 1
  Interested    MessageID = 2
  NotInterested MessageID = 3
  Have          MessageID = 4
  Bitfield      MessageID = 5
  Request       MessageID = 6
  Piece         MessageID = 7
  Cancel        MessageID = 8
  Port          MessageID = 9
)

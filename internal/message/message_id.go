package message

type MessageID int

const (
  ChokeID         MessageID = 0
  UnchokeID       MessageID = 1
  InterestedID    MessageID = 2
  NotInterestedID MessageID = 3
  HaveID          MessageID = 4
  BitfieldID      MessageID = 5
  RequestID       MessageID = 6
  PieceID         MessageID = 7
  CancelID        MessageID = 8
  PortID          MessageID = 9
)

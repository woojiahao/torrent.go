package p2p

type messageID int

const (
  choke         messageID = 0
  unchoke       messageID = 1
  interested    messageID = 2
  notInterested messageID = 3
  have          messageID = 4
  bitfield      messageID = 5
  request       messageID = 6
  piece         messageID = 7
  cancel        messageID = 8
  port          messageID = 9
)

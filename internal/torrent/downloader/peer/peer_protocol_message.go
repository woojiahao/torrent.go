package peer

type peerMessage struct {
  length  string
  message string
  payload string
}

func (pm *peerMessage) serialize() []byte {

  return []byte{}
}

func deserialize(b []byte) *peerMessage {

}


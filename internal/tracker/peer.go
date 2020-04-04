package tracker

import (
  "net"
  "strconv"
)

type Peer struct {
  ip   string
  port int
}

func (p *Peer) Address() string {
  return net.JoinHostPort(p.ip, strconv.Itoa(p.port))
}

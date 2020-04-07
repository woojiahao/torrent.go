package tracker

import (
  "net"
  "strconv"
)

type Peer struct {
  ip   net.IP
  port uint16
}

func (p *Peer) Address() string {
  return net.JoinHostPort(p.ip.String(), strconv.Itoa(int(p.port)))
}

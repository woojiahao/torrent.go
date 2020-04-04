package utility

import (
  "errors"
  "github.com/woojiahao/torrent.go/internal/downloader/p2p"
  "net"
)

type TCPConn struct {
  Address string
  Conn    net.Conn
}

// Establishes a TCP connection with a given IP address.
// The connection will timeout after a given amount of seconds.
func TCP(address string, timeout int) (*TCPConn, error) {
  c, err := net.DialTimeout("tcp", address, ToSeconds(timeout))
  return &TCPConn{address, c}, err
}

// Write to a TCP connection
func (c *TCPConn) Send(data []byte) error {
  n, err := c.Conn.Write(data)
  if n != len(data) {
    return errors.New("unsuccessful; bytes written not equal to the bytes sent")
  }

  return err
}

// Writes a message to the TCP connection
func (c *TCPConn) SendMessage(id p2p.MessageID, payload ...byte) error {
  m := p2p.Message{
    MessageID: id,
    Payload:   payload,
  }
  err := c.Send(m.Serialize())
  return err
}

// Read from a TCP connection
func (c *TCPConn) Receive(bufSize int) ([]byte, error) {
  buf := make([]byte, bufSize)
  _, err := c.Conn.Read(buf)
  return buf, err
}

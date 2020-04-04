package utility

import (
  "errors"
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

// Read from a TCP connection
func (c *TCPConn) Receive(bufSize int) ([]byte, error) {
  buf := make([]byte, bufSize)
  _, err := c.Conn.Read(buf)
  return buf, err
}

func (c *TCPConn)

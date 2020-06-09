package connection

import (
  "net"
  "time"
)

type Connection struct {
  Address string
  Conn    net.Conn
}

// Establishes a TCP connection with a given IP address.
// The connection will timeout after a given amount of seconds.
func TCP(address string, timeout int) (*Connection, error) {
  c, err := net.DialTimeout("tcp", address, time.Duration(timeout)*time.Second)
  return &Connection{address, c}, err
}

// Write to a TCP connection
func (c *Connection) Send(data []byte) error {
  _, err := c.Conn.Write(data)
  return err
}

// Read from a TCP connection
func (c *Connection) Receive(size int) ([]byte, error) {
  buf := make([]byte, size)
  _, err := c.Conn.Read(buf)
  return buf, err
}

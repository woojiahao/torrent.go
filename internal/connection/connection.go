package connection

import (
  "errors"
  "io/ioutil"
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
  n, err := c.Conn.Write(data)
  if n != len(data) {
    return errors.New("unsuccessful; bytes written not equal to the bytes sent")
  }

  return err
}

// Read from a TCP connection
func (c *Connection) Receive() ([]byte, error) {
  buf, err := ioutil.ReadAll(c.Conn)
  return buf, err
}

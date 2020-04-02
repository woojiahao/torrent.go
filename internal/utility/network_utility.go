package utility

import (
  "net"
  . "net/http"
)

type QueryParameters map[string]string

func GET(URL string, parameters *QueryParameters) *Response {
  client := Client{}
  req, err := NewRequest("GET", URL, nil)
  LogCheck(err)

  q := req.URL.Query()
  for key, value := range *parameters {
    q.Add(key, value)
  }

  req.URL.RawQuery = q.Encode()

  resp, err := client.Do(req)
  LogCheck(err)

  return resp
}

// Establishes a TCP connection with a given IP address.
// The connection will timeout after a given amount of seconds.
func TCP(address string, timeout int) (c net.Conn, err error) {
  c, err = net.DialTimeout("tcp", address, ToSeconds(timeout))
  return c, err
}

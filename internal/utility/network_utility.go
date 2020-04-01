package utility

import (
  . "net/http"
)

type QueryParameters map[string]string

func GET(URL string, parameters QueryParameters) *Response {
  client := Client{}
  req, err := NewRequest("GET", URL, nil)
  Check(err)

  q := req.URL.Query()
  for key, value := range parameters {
    q.Add(key, value)
  }

  req.URL.RawQuery = q.Encode()

  resp, err := client.Do(req)
  Check(err)

  return resp
}

package torrent_file

import "fmt"

type (
  fileError struct {
    filename string
    reason string
  }
)

func (err *fileError) Error() string {
  return fmt.Sprintf("file: %s %s", err.filename, err.reason)
}

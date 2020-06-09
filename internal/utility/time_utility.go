package utility

import "time"

func ToSeconds(seconds int) time.Duration {
  return time.Duration(seconds) * time.Second
}

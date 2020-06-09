package bencoding

import "fmt"

type (
  typeConversionError struct {
    targetType string
  }

  decodeError struct {
    reason string
  }
)

func (err *typeConversionError) Error() string {
  return fmt.Sprintf("failed to convert TType to %s", err.targetType)
}

func (err *decodeError) Error() string {
  return fmt.Sprintf("bencoding decoding failed due to %s", err.reason)
}

package utility

import (
  "fmt"
  "path/filepath"
)

// Check if a file's file type is within a specified selection
func IsFileType(filename string, extensions ...string) bool {
  ext := filepath.Ext(filename)
  for _, e := range extensions {
    if ext == fmt.Sprintf(".%s", e) {
      return true
    }
  }
  return false
}

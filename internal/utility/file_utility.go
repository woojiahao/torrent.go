package utility

import (
  "fmt"
  "io/ioutil"
  "os"
  "path/filepath"
)

// Reads a file's content
func ReadFileContents(filename string) string {
  data, err := ioutil.ReadFile(filename)
  Check(err)
  return string(data)
}

// Check if a given file exists
func Exists(filename string) bool {
  _, err := os.Stat(filename)
  return os.IsNotExist(err)
}

// Check if a given filename is a directory
func IsDir(filename string) bool {
  file, err := os.Stat(filename)
  Check(err)
  return file.IsDir()
}

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

package utility

import "io/ioutil"

// Reads a file's content
func ReadFileContents(filename string) string {
  data, err := ioutil.ReadFile(filename)
  Check(err)
  return string(data)
}

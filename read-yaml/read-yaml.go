package main

import (
  "fmt"
  "os"
  "bytes"
  "io/ioutil"
  //"gopkg.in/yaml.v2"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readConfig() string {
  var path bytes.Buffer
  homeDir := os.Getenv("HOME")
  path.WriteString(homeDir)
  path.WriteString("/.config/search-service/config.yml")
  dat, err := ioutil.ReadFile(path.String())
  check(err)
  return string(dat)
}

func main() {
  config := readConfig()
  fmt.Println(config)
}

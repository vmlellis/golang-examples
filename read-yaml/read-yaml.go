package main

import (
  "fmt"
  "os"
  "bytes"
  "io/ioutil"
  "gopkg.in/yaml.v2"
  "github.com/davecgh/go-spew/spew"
)

type Host struct {
  Host string     `yaml:"host"`
  Port int        `yaml:"port"`
  User string     `yaml:"user"`
  Password string `yaml:"password"`
  scheme string   `yaml:"scheme"`
}

type Config struct {
    Hosts []Host `yaml:"hosts"`
}


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
  configText := readConfig()
  fmt.Println(configText)
  config := Config{}
  err := yaml.Unmarshal([]byte(configText), &config)
  check(err)
  spew.Dump(config)
  fmt.Println(config)
  fmt.Println(config.Hosts)
  fmt.Println(config.Hosts[0].Host)
  fmt.Println(config.Hosts[1].Host)

  m := make(map[string][]map[string]interface{})
  err = yaml.Unmarshal([]byte(configText), &m)
  check(err)
  spew.Dump(m)
  hosts := m["hosts"]

  for _, host := range hosts {
    for k, v := range host {
      fmt.Printf("%s -> %s\n", k, v)
    }
  }
}

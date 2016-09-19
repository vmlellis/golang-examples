package main

import (
  "fmt"
  "flag"
  "strings"
)

func main() {
  argKeys := flag.String("keys", "", "keys to test")
  flag.Parse()
  printKeys(*argKeys)

}

func printKeys(str string) {
  values := strings.Split(str, ",")
  for _, value := range values {
    fmt.Println("key:", value)
  }
}

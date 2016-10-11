package main

import (
  "gopkg.in/olivere/elastic.v3"
  "github.com/davecgh/go-spew/spew"
)

func setupClient() (client *elastic.Client) {
  client, err := elastic.NewClient()
  if err != nil {
    // Handle error
    panic(err)
  }
  return client
}

func dumpMultipleSort() {
  query := elastic.NewSearchSource().
    Sort("reference", true).
    Sort("action", true).
    Size(100)
  spew.Dump(query.Source())
}

func main() {
  dumpMultipleSort()
}

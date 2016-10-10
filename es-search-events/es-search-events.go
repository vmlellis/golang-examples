package main

import (
  "io"
  "gopkg.in/olivere/elastic.v3"
  "github.com/davecgh/go-spew/spew"
)

const (
  indexName  = "index"
  typeName   = "type"
  parentName = "parent"
)

func setupClient() (client *elastic.Client) {
  client, err := elastic.NewClient()
  if err != nil {
    // Handle error
    panic(err)
  }
  return client
}

func dumpChildren(client *elastic.Client) {
  termQuery := elastic.NewTermQuery("_parent", parentName)
  scroll := client.Scroll(indexName).
    Type(typeName).
    Query(termQuery).
    Sort("action", true).
    Size(100)
  for {
    results, err := scroll.Do()
    if err == io.EOF {
      break // all results retrieved
    }
    if err != nil {
      break // something went wrong
    }

    for _, hit := range results.Hits.Hits {
      spew.Dump(hit.Source)
    }
  }
}

func main() {
  client := setupClient()
  dumpChildren(client)
}

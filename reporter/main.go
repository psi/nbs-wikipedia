package main

import (
  //"bufio"
  "fmt"
  //"log"
  //"os"
  //"regexp"
  //"strconv"
  //"strings"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type Page struct {
  Language string
  Title    string
  Views    int
}

func main() {
  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  defer session.Close()

  pages := session.DB("wikipedia").C("pages")

  var languages []string
  err = pages.Find(bson.M{}).Distinct("language", &languages)

  if err != nil {
    panic(err)
  }

  fmt.Println("Found ", len(languages), " languages")
}

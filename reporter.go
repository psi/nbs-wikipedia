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

  for i := 0; i < len(languages); i++ {
    var topPages []Page
    pages.Find(bson.M{ "language": languages[i] }).Sort("-views").Limit(5).All(&topPages)

    fmt.Println(languages[i], ":")
    for j := 0; j < len(topPages); j++ {
      fmt.Println(topPages[j].Title, ",", topPages[j].Views)
    }
  }

  fmt.Println("Found ", len(languages), " languages")
}

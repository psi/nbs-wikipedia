package main

import (
  "bufio"
  //"fmt"
  "log"
  "os"
  "regexp"
  "strconv"
  "strings"
  "gopkg.in/mgo.v2"
  //"gopkg.in/mgo.v2/bson"
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

  file, err := os.Open("/tmp/pagecounts.csv")
  if err != nil {
    panic(err)
  }

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    fields := strings.Split(scanner.Text(), " ")

    language := fields[0]
    title    := fields[1]

    // Skip titles that contain prefixes like "Special:", "User:", etc.
    matched, err := regexp.MatchString(":", title)
    if err != nil {
      panic(err)
    }

    if !matched {
      views, err := strconv.Atoi(fields[3])
      if err != nil {
        panic(err)
      }

      err = pages.Insert(&Page{language, title, views})

      if err != nil {
        log.Fatal(err)
      }
    }
  }
}

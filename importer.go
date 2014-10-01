package main

import (
	"bufio"
	//"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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
		title := fields[1]

		// We only want pages from Wikipedia, not other projects
		isNotWikipediaPage, err := regexp.MatchString("\\.", language)
		if err != nil {
			panic(err)
		}

		// Skip titles that contain prefixes like "Special:", "User:", etc.
		isNonStandardPage, err := regexp.MatchString(":", title)
		if err != nil {
			panic(err)
		}

		if !isNotWikipediaPage && !isNonStandardPage {
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

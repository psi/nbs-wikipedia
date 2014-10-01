package main

import (
	//"bufio"
	//"fmt"
	//"log"
	//"os"
	//"regexp"
	//"strconv"
	//"strings"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"sync"
)

type Page struct {
	Language string
	Title    string
	Views    int
}

const databaseName = "wikipedia"
const sourceCollectionName = "pages"

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	pages := session.DB(databaseName).C(sourceCollectionName)

	var languages []string
	err = pages.Find(bson.M{}).Distinct("language", &languages)

	if err != nil {
		panic(err)
	}

	var waitGroup sync.WaitGroup

	for i := 0; i < len(languages); i++ {
		waitGroup.Add(1)
		go processLanguage(languages[i], &waitGroup, session)
	}

	waitGroup.Wait()
}

func processLanguage(language string, waitGroup *sync.WaitGroup, mongoSession *mgo.Session) {
	defer waitGroup.Done()

	session := mongoSession.Copy()
	defer session.Close()

	var topPages []Page

	pages := mongoSession.DB(databaseName).C(sourceCollectionName)
	pages.Find(bson.M{"language": language}).Sort("-views").Limit(10).All(&topPages)

	targetCollection := session.DB(databaseName).C("top_pages_" + language)
	targetCollection.Insert(topPages)
}

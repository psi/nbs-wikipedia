package main

import (
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

  // Find all of the languages we have records for
	var languages []string
	err = pages.Find(bson.M{}).Distinct("language", &languages)

	if err != nil {
		panic(err)
	}

  // Launch a bunch of goroutines to calculate the top 10 pages per language.
  // Ideally, these would be load-balanced to multiple MongoDB replica set read
  // slaves, either by mgo's connection pooling or sitting behind HAProxy.
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

  // If we were load-balancing reads behind HAProxy, we'd need to take care
  // here to write back to the master MongoDB instance.
	targetCollection := session.DB(databaseName).C("top_pages_" + language)
	targetCollection.Insert(topPages)
}

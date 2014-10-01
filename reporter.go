package main

import (
	//"bufio"
	"fmt"
	//"log"
	//"os"
	//"regexp"
	//"strconv"
  //"strings"
  //"sync"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

	for i := 0; i < len(languages); i++ {
		var topPages []Page
		pages.Find(bson.M{"language": languages[i]}).Sort("-views").Limit(10).All(&topPages)

    targetCollection := session.DB(databaseName).C("top_pages_" + languages[i])

    targetCollection.Insert(topPages)

		//fmt.Println(languages[i], ":")
		//for j := 0; j < len(topPages); j++ {
		//	fmt.Println(topPages[j].Title, ",", topPages[j].Views)
		//}

		//fmt.Println()
	}
}

package utils

import (
	"fmt"
	_ "github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func GetBookByName(name string) book {
	var books book
	session := connect()
	defer session.Close()
	err := session.DB("bookinfo").C("books").Find(bson.M{"Title":name}).One(&books)
	if err != nil {
		log.Println(err)
		return book{}
	}

	fmt.Println(books)
	return books
}

func connect() *mgo.Session {
	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println("session err:", err)
		return nil
	}
	return session
}

func GetBooks()[]book{
	var books []book
	session := connect()
	defer session.Close()
	err := session.DB("bookinfo").C("books").Find(bson.M{}).All(&books)
	if err != nil {
		log.Println(err)
		return nil
	}

	for _,v:=range books{
		fmt.Println(v.Title)
	}

	//fmt.Println(books)
	return books
}

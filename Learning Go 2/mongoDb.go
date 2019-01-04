package main

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	FirstName string     `json:"first_name" bson:"first_name,omitempty"`
	LastName  string     `json:"last_name" bson:"last_name,omitempty"`
	Grade     string     `json:"grade" bson:"grade,omitempty"`
}

func main() {

	session, err := mgo.Dial("localhost")

	if (err != nil) {
		fmt.Println("error in connecting !", err)
		return
	}
	defer session.Close()
	fmt.Println("connected", session)

	dbs, _ := session.DatabaseNames()
	fmt.Println(dbs)

	collections, _ := session.DB("sampleDb").CollectionNames()
	fmt.Println(collections)

	c := session.DB("sampleDb").C("customers")

	//fmt.Println(c.Find(bson.M{"_id":}))

	//Insert
	//rec1 := Person{Id:bson.NewObjectId(),FirstName:"Harsh", LastName:"Sidna"}
	//c.Insert(&rec1)

	fmt.Println(c.Count())

	//Query All
	var results []Person
	c.Find(nil).All(&results)
	for _, v := range results {
		fmt.Println("All Records :" , v)
	}

	//Query One
	//var record  Person
	//c.Find(bson.M{"last_name": "Patel"}).One(&record)
	//fmt.Println(record)

	//Query by Id
	//var record1 Person
	//id:= bson.ObjectIdHex("5b50326bf57d8774419c0405")
	//c.FindId(id).One(&record1)
	//fmt.Println(record1)

	//Remove by Id
	//c.RemoveId(id)

	//Update
	//c.Update(bson.M{"first_name":"Shubham"},bson.M{"last_name":"Gargg"})                       //deletes empty fields
	//c.Update(bson.M{"first_name":"Shubham"},bson.M{"$set":bson.M{"last_name":"Gargg"}})          //retains empty fields

	//Update by Id
	//updateId:= bson.ObjectIdHex("5b50326bf57d8774419c0405")
	//c.UpdateId(updateId,bson.M{"first_name":"Shubham","last_name":"Garg"})

}

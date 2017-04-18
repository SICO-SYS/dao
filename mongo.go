/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	. "github.com/SiCo-DevOps/log"
)

type User struct {
	Id     string "_id"
	Key    string "key"
	Secret string "secret"
}

var (
	MgoConn *mgo.Session
)

func init() {
	MgoConn, err := mgo.Dial(config.Mongo.User.Addr)
	if err != nil {
		LogProduce("error", "Mongo connect failed")
	}
	MgoConn.SetPoolLimit(10)
	defer MgoConn.Close()
	result := User{}
	c := MgoConn.Clone()
	s := c.DB("user").C("keypair")
	// s.EnsureIndex(index)
	s.Find(bson.M{"key": "123456"}).One(&result)
	// defer s.Close()
	c.Close()
	// c := s.DB("SiCo").C("user")
	// c.Find("123456")
}

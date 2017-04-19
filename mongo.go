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

type UserKeypair struct {
	Id     bson.ObjectId "_id"
	Key    string        "key"
	Secret string        "secret"
}

var (
	MgoConn          *mgo.Session
	UserKeypairIndex = mgo.Index{
		Key:        []string{"key"},
		Unique:     true,
		DropDups:   true,
		Background: true,
	}
)

// func EnsureIndex(c *mgo.Collection, index mgo.Index) bool {
// 	err = c.EnsureIndex(index)
// 	if err != nil {
// 		LogErrMsg(22, "dao.EnsureIndex")
// 		return false
// 	}
// 	return true
// }

func (u *UserKeypair) Insert(k string, s string) bool {
	user := &UserKeypair{bson.NewObjectId(), k, s}
	c := MgoConn.Clone()
	defer c.Close()
	err = c.DB("SiCo").C("user.keypair").Insert(user)
	if err != nil {
		LogErrMsg(21, "dao.UserKeypair.Insert")
		return false
	}
	return true
}

func init() {
	MgoConn, err := mgo.Dial(config.Mongo.UserKeypair.Addr)
	if err != nil {
		LogErrMsg(2, "dao.init")
	} else {
		MgoConn.SetPoolLimit(10)
	}
	MgoConn.DB("SiCo").C("user.keypair").EnsureIndex(UserKeypairIndex)
}

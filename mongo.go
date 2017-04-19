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
	MgoUserConn, mgoerr = mgo.Dial(config.Mongo.UserKeypair.Addr)
	UserKeypairIndex    = mgo.Index{
		Key:        []string{"key"},
		Unique:     true,
		DropDups:   true,
		Background: true,
	}
)

func (u *UserKeypair) Insert() bool {
	defer func() {
		if recover() != nil {
			LogErrMsg(2, "dao.mongo.Insert")
		}
	}()
	user := &UserKeypair{bson.NewObjectId(), u.Key, u.Secret}
	c := MgoUserConn.Clone()
	defer c.Close()
	err = c.DB("SiCo").C("user.keypair").Insert(user)
	if err != nil {
		LogErrMsg(21, "dao.UserKeypair.Insert")
		return false
	}
	return true
}

func init() {
	if mgoerr != nil {
		LogErrMsg(2, "dao.init")
	} else {
		MgoUserConn.SetPoolLimit(100)
	}
	MgoUserConn.DB("SiCo").C("user.keypair").EnsureIndex(UserKeypairIndex)
}

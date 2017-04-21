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

var (
	MgoUserConn, mgoerr = mgo.Dial(config.Mongo.User.Addr)
	UserKeyIndex        = mgo.Index{
		Key:        []string{"key"},
		Unique:     true,
		DropDups:   true,
		Background: true,
	}
)

func Mgo_Insert(v interface{}, c string) bool {
	defer func() {
		if recover() != nil {
			LogErrMsg(2, "dao.Mgo_Insert")
		}
	}()
	conn := MgoUserConn.Clone()
	defer conn.Close()
	err = conn.DB("SiCo").C(c).Insert(v)
	if err != nil {
		LogErrMsg(21, "dao.Mgo_Insert")
		return false
	}
	return true
}

func Mgo_Find(k string, s string) map[string]interface{} {
	return bson.M{k: s}
}

func init() {
	if mgoerr != nil {
		LogErrMsg(2, "dao.init")
	} else {
		MgoUserConn.SetPoolLimit(100)
	}
	MgoUserConn.DB("SiCo").C("user.keypair").EnsureIndex(UserKeyIndex)
	MgoUserConn.DB("SiCo").C("user.auth").EnsureIndex(UserKeyIndex)
}

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
	MongoIdIndex        = mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
	}
	MongoIDNameIndex = mgo.Index{
		Key:        []string{"id", "name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
	}
)

func Mgo_Insert(v interface{}, c string) bool {
	defer func() {
		recover()
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

func AAA_ensureIndexes() {
	MgoUserConn.DB("SiCo").C("user.keypair").EnsureIndex(MongoIdIndex)
	MgoUserConn.DB("SiCo").C("user.auth").EnsureIndex(MongoIdIndex)
	MgoUserConn.DB("SiCo").C("user.cloud.aws").EnsureIndex(MongoIDNameIndex)
	MgoUserConn.DB("SiCo").C("user.cloud.aliyun").EnsureIndex(MongoIDNameIndex)
	MgoUserConn.DB("SiCo").C("user.cloud.qcloud").EnsureIndex(MongoIDNameIndex)

}

func init() {
	defer func() {
		recover()
		if recover() != nil {
			LogProduce("error", "Maybe mongo connection failed")
		}
	}()
	if mgoerr != nil {
		LogErrMsg(2, "dao.init")
	} else {
		MgoUserConn.SetPoolLimit(100)
	}
}

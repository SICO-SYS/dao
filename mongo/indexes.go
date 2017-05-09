/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package mongo

import (
	"gopkg.in/mgo.v2"
)

var (
	MongoIdIndex = mgo.Index{
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

func AAA_ensureIndexes() {
	MgoUserConn.DB("SiCo").C("user.keypair").EnsureIndex(MongoIdIndex)
	MgoUserConn.DB("SiCo").C("user.auth").EnsureIndex(MongoIdIndex)
	MgoUserConn.DB("SiCo").C("user.cloud.aws").EnsureIndex(MongoIDNameIndex)
	MgoUserConn.DB("SiCo").C("user.cloud.aliyun").EnsureIndex(MongoIDNameIndex)
	MgoUserConn.DB("SiCo").C("user.cloud.qcloud").EnsureIndex(MongoIDNameIndex)
}

func Cloud_ensureIndexes() {
	MgoUserConn.DB("SiCo").C("cloud.aws.template").EnsureIndex(MongoIDNameIndex)
	MgoUserConn.DB("SiCo").C("cloud.aliyun.template").EnsureIndex(MongoIDNameIndex)
	MgoUserConn.DB("SiCo").C("cloud.qcloud.template").EnsureIndex(MongoIDNameIndex)
}

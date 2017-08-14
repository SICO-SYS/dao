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
	MongoNameIndex = mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
	}
)

func AAAEnsureIndexes() {
	MgoUserConn.DB("SiCo").C("user.token").EnsureIndex(MongoIdIndex)
	MgoUserConn.DB("SiCo").C("user.policy").EnsureIndex(MongoIdIndex)
}

func CloudEnsureIndexes() {
	MgoCloudConn.DB("SiCo").C("cloud.token.aws").EnsureIndex(MongoIDNameIndex)
	MgoCloudConn.DB("SiCo").C("cloud.token.aliyun").EnsureIndex(MongoIDNameIndex)
	MgoCloudConn.DB("SiCo").C("cloud.token.qcloud").EnsureIndex(MongoIDNameIndex)

}

func AssetEnsureIndexes(id string) {
	// MgoAssetConn.DB("SiCo").C("asset." + id + "." + cloud).EnsureIndex(MongoIdIndex)
	MgoAssetConn.DB("SiCo").C("template." + id).EnsureIndex(MongoNameIndex)
}

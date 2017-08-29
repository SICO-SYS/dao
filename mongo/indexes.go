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
	MgoUserConn.DB(databaseName()).C(CollectionUserTokenName()).EnsureIndex(MongoIdIndex)
	MgoUserConn.DB(databaseName()).C(CollectionUserPolicyName()).EnsureIndex(MongoIdIndex)
}

func CloudEnsureIndexes() {
	MgoCloudConn.DB(databaseName()).C(CollectionCloudTokenName("aws")).EnsureIndex(MongoIDNameIndex)
	MgoCloudConn.DB(databaseName()).C(CollectionCloudTokenName("aliyun")).EnsureIndex(MongoIDNameIndex)
	MgoCloudConn.DB(databaseName()).C(CollectionCloudTokenName("qcloud")).EnsureIndex(MongoIDNameIndex)

}

func TemplateEnsureIndexes(id string) {
	MgoAssetConn.DB(databaseName()).C(CollectionTemplateName(id)).EnsureIndex(MongoNameIndex)
}

func AssetEnsureIndexes(cloud, id string) {
	MgoAssetConn.DB(databaseName()).C(CollectionAssetCloudName(cloud, id)).EnsureIndex(MongoNameIndex)

}

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
	IdIndex = mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
	}
	IDNameIndex = mgo.Index{
		Key:        []string{"id", "name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
	}
	NameIndex = mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
	}
)

func AAAEnsureIndexes() {
	defer func() {
		recover()
	}()
	conn := UserConn.Clone()
	defer conn.Close()
	conn.DB(databaseName()).C(CollectionUserTokenName()).EnsureIndex(IdIndex)
	conn.DB(databaseName()).C(CollectionUserPolicyName()).EnsureIndex(IdIndex)
}

func CloudEnsureIndexes() {
	defer func() {
		recover()
	}()
	conn := CloudConn.Clone()
	defer conn.Close()
	conn.DB(databaseName()).C(CollectionCloudTokenName("aws")).EnsureIndex(IDNameIndex)
	conn.DB(databaseName()).C(CollectionCloudTokenName("aliyun")).EnsureIndex(IDNameIndex)
	conn.DB(databaseName()).C(CollectionCloudTokenName("qcloud")).EnsureIndex(IDNameIndex)

}

func TemplateEnsureIndexes(id string) {
	defer func() {
		recover()
	}()
	conn := AssetConn.Clone()
	defer conn.Close()
	conn.DB(databaseName()).C(CollectionTemplateName(id)).EnsureIndex(NameIndex)
}

func AssetEnsureIndexes(cloud, id string) {
	defer func() {
		recover()
	}()
	conn := AssetConn.Clone()
	defer conn.Close()
	conn.DB(databaseName()).C(CollectionAssetCloudName(cloud, id)).EnsureIndex(NameIndex)

}

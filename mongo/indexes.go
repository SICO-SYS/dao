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

func AAAEnsureIndexes(m *mgo.Session) {
	defer func() {
		recover()
	}()
	conn := m.Clone()
	defer conn.Close()
	conn.DB(databaseName()).C(CollectionUserTokenName()).EnsureIndex(IdIndex)
	conn.DB(databaseName()).C(CollectionUserPolicyName()).EnsureIndex(IdIndex)
}

func CloudEnsureIndexes(m *mgo.Session) {
	defer func() {
		recover()
	}()
	conn := m.Clone()
	defer conn.Close()
	conn.DB(databaseName()).C(CollectionCloudTokenName("aws")).EnsureIndex(IDNameIndex)
	conn.DB(databaseName()).C(CollectionCloudTokenName("aliyun")).EnsureIndex(IDNameIndex)
	conn.DB(databaseName()).C(CollectionCloudTokenName("qcloud")).EnsureIndex(IDNameIndex)

}

func TemplateEnsureIndexes(m *mgo.Session, id string) {
	defer func() {
		recover()
	}()
	conn := m.Clone()
	defer conn.Close()
	conn.DB(databaseName()).C(CollectionTemplateName(id)).EnsureIndex(NameIndex)
}

func AssetEnsureIndexes(m *mgo.Session, cloud, id string) {
	defer func() {
		recover()
	}()
	conn := m.Clone()
	defer conn.Close()
	conn.DB(databaseName()).C(CollectionAssetCloudName(cloud, id)).EnsureIndex(NameIndex)

}

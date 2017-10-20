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

func AAAEnsureIndexes(m *mgo.Session) (err error) {
	conn := m.Clone()
	err = m.Ping()
	if err != nil {
		return err
	}
	defer conn.Close()
	err = conn.DB(databaseName()).C(CollectionUserTokenName()).EnsureIndex(IdIndex)
	if err != nil {
		return err
	}
	err = conn.DB(databaseName()).C(CollectionUserPolicyName()).EnsureIndex(IdIndex)
	if err != nil {
		return err
	}
	return nil
}

func CloudEnsureIndexes(m *mgo.Session) (err error) {
	conn := m.Clone()
	err = m.Ping()
	if err != nil {
		return err
	}
	defer conn.Close()
	err = conn.DB(databaseName()).C(CollectionCloudTokenName("aws")).EnsureIndex(IDNameIndex)
	if err != nil {
		return err
	}
	err = conn.DB(databaseName()).C(CollectionCloudTokenName("aliyun")).EnsureIndex(IDNameIndex)
	if err != nil {
		return err
	}
	err = conn.DB(databaseName()).C(CollectionCloudTokenName("qcloud")).EnsureIndex(IDNameIndex)
	if err != nil {
		return err
	}
	return nil
}

func TemplateEnsureIndexes(m *mgo.Session, id string) (err error) {
	conn := m.Clone()
	err = m.Ping()
	if err != nil {
		return err
	}
	defer conn.Close()
	err = conn.DB(databaseName()).C(CollectionTemplateName(id)).EnsureIndex(NameIndex)
	if err != nil {
		return err
	}
	return nil
}

func AssetEnsureIndexes(m *mgo.Session, cloud, id string) (err error) {
	conn := m.Clone()
	err = m.Ping()
	if err != nil {
		return err
	}
	defer conn.Close()
	err = conn.DB(databaseName()).C(CollectionAssetCloudName(cloud, id)).EnsureIndex(NameIndex)
	if err != nil {
		return err
	}
	return nil
}

func HookEnsureIndexes(m *mgo.Session) (err error) {
	conn := m.Clone()
	err = m.Ping()
	if err != nil {
		return err
	}
	defer conn.Close()
	err = conn.DB(databaseName()).C(CollectionHookName()).EnsureIndex(IdIndex)
	if err != nil {
		return err
	}
	err = conn.DB(databaseName()).C(CollectionHookName()).EnsureIndex(NameIndex)
	if err != nil {
		return err
	}
	return nil
}

/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package mongo

import (
	"github.com/getsentry/raven-go"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Querys map[string]string

func Insert(mgoconn *mgo.Session, v interface{}, c string) bool {
	defer func() {
		recover()
	}()
	conn := mgoconn.Clone()
	defer conn.Close()
	err := conn.DB(databaseName()).C(c).Insert(v)
	if err != nil {
		raven.CaptureError(err, nil)
		return false
	}
	return true
}

func Find(k string, s string) map[string]interface{} {
	return bson.M{k: s}
}

func (m Querys) FindOne(mgoconn *mgo.Session, c string) bson.M {
	defer func() {
		recover()
	}()
	data, _ := bson.Marshal(m)
	query := bson.M{}
	bson.Unmarshal(data, query)
	conn := mgoconn.Clone()
	defer conn.Close()
	b := bson.M{}
	conn.DB(databaseName()).C(c).Find(query).One(b)
	return b
}

func FindOne(mgoconn *mgo.Session, q Querys) (m map[string]interface{}) {
	conn := mgoconn.Clone()
	defer conn.Close()
	conn.DB("SiCo").C("user.token").Find(q).One(&m)
	return m
}

func Remove(mgoconn *mgo.Session, c string) bool {
	defer func() {
		recover()
	}()
	conn := mgoconn.Clone()
	defer conn.Close()
	err := conn.DB(databaseName()).C(c).Remove(bson.M{})
	if err != nil {
		raven.CaptureError(err, nil)
		return false
	}
	return true
}

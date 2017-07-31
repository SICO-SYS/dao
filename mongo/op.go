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

type MgoQuerys map[string]string

func MgoInsert(mgoconn *mgo.Session, v interface{}, c string) bool {
	defer func() {
		recover()
		if recover() != nil {
			raven.CaptureMessage("dao.mongo.Mgo_Insert", nil)
		}
	}()
	conn := mgoconn.Clone()
	defer conn.Close()
	err := conn.DB("SiCo").C(c).Insert(v)
	if err != nil {
		raven.CaptureError(err, nil)
		return false
	}
	return true
}

func MgoFind(k string, s string) map[string]interface{} {
	return bson.M{k: s}
}

func (m MgoQuerys) MgoFindOne(mgoconn *mgo.Session, c string) bson.M {
	defer func() {
		recover()
		if recover() != nil {
			raven.CaptureMessage("dao.mongo.MgoFindOne", nil)
		}
	}()
	data, _ := bson.Marshal(m)
	query := bson.M{}
	bson.Unmarshal(data, query)
	conn := mgoconn.Clone()
	defer conn.Close()
	b := bson.M{}
	conn.DB("SiCo").C(c).Find(query).One(b)
	return b
}

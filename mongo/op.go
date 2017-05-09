/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	. "github.com/SiCo-DevOps/log"
)

type Mgo_Querys map[string]string

func Mgo_Insert(mgoconn *mgo.Session, v interface{}, c string) bool {
	defer func() {
		recover()
		if recover() != nil {
			LogErrMsg(2, "dao.Mgo_Insert")
		}
	}()
	conn := mgoconn.Clone()
	defer conn.Close()
	err := conn.DB("SiCo").C(c).Insert(v)
	if err != nil {
		LogErrMsg(21, "dao.Mgo_Insert")
		return false
	}
	return true
}

func Mgo_Find(k string, s string) map[string]interface{} {
	return bson.M{k: s}
}

func (m Mgo_Querys) Mgo_FindsOne(mgoconn *mgo.Session, c string) bson.M {
	defer func() {
		recover()
		if recover() != nil {
			LogErrMsg(2, "dao.Mgo_FindsOne")
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

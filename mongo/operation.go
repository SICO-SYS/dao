/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package mongo

import (
	"gopkg.in/mgo.v2"
)

type Querys map[string]string

func Insert(mgoconn *mgo.Session, c string, v interface{}) (err error) {
	conn := mgoconn.Clone()
	err = mgoconn.Ping()
	if err != nil {
		return err
	}
	defer conn.Close()
	err = conn.DB(databaseName()).C(c).Insert(v)
	return err
}

func FindOne(mgoconn *mgo.Session, c string, q map[string]string) (m map[string]interface{}, err error) {
	conn := mgoconn.Clone()
	err = mgoconn.Ping()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	err = conn.DB(databaseName()).C(c).Find(q).One(&m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func Remove(mgoconn *mgo.Session, c string, q map[string]string) (err error) {
	conn := mgoconn.Clone()
	err = mgoconn.Ping()
	if err != nil {
		return err
	}
	defer conn.Close()
	err = conn.DB(databaseName()).C(c).Remove(q)
	if err != nil {
		return err
	}
	return nil
}

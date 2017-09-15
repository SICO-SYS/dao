/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package mongo

import (
	"gopkg.in/mgo.v2"
)

func databaseName() string {
	return "SiCo"
}

func Dial(address, username, password) *mgo.Session {
	return mgo.Dial(address)
}

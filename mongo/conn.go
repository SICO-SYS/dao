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

var (
	TestingConn, TestingErr = mgo.Dial(config.Mongo.Testing.Address)
	UserConn, UserErr       = mgo.Dial(config.Mongo.User.Address)
	CloudConn, CloudErr     = mgo.Dial(config.Mongo.Cloud.Address)
	AssetConn, AssetErr     = mgo.Dial(config.Mongo.Asset.Address)
)

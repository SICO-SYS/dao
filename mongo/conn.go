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
	MgoTestingConn, mgoTestingErr = mgo.Dial(config.Mongo.Testing.Address)
	MgoUserConn, mgoUserErr       = mgo.Dial(config.Mongo.User.Address)
	MgoCloudConn, mgoCloudErr     = mgo.Dial(config.Mongo.Cloud.Address)
	MgoAssetConn, mgoAssetErr     = mgo.Dial(config.Mongo.Asset.Address)
)

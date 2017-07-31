/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package mongo

import (
	"github.com/getsentry/raven-go"
	"gopkg.in/mgo.v2"
)

var (
	MgoDefaultConn, mgoDefaultErr = mgo.Dial(config.Mongo.Default.Addr)
	MgoUserConn, mgoUserErr       = mgo.Dial(config.Mongo.User.Addr)
	MgoCloudConn, mgoCloudErr     = mgo.Dial(config.Mongo.Cloud.Addr)
	MgoAssetConn, mgoAssetErr     = mgo.Dial(config.Mongo.Asset.Addr)
)

func init() {
	defer func() {
		recover()
		if recover() != nil {
			raven.CaptureMessage("dao.mongo.init", nil)
		}
	}()

	MgoDefaultConn.SetPoolLimit(10)
	MgoUserConn.SetPoolLimit(10)
	MgoCloudConn.SetPoolLimit(10)
	MgoAssetConn.SetPoolLimit(10)

}

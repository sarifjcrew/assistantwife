package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type ShopModel struct {
	Id         bson.ObjectId `bson:"_id",json:"Id"`
	Tanggal    time.Time     `json:"Tanggal"`
	Keterangan string        `json:"Keterangan"`
	Harga      int64         `json:"Harga"`
}

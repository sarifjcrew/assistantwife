package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type DebtModel struct {
	Id         bson.ObjectId `bson:"_id",json:"Id"`
	Tanggal    time.Time     `json:"Tanggal"`
	Keterangan string        `json:"Keterangan"`
	Debt       int64         `json:"Debt"`
}

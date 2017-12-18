package models

import (
	"gopkg.in/mgo.v2/bson"
)

type SaldoModel struct {
	Id    bson.ObjectId `bson:"_id",json:"_id"`
	Saldo int64         `json:Saldo`
}

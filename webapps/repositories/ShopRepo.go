package repositories

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"pribadi/assistantwife/webapps/models"
)

type ShopRepo struct {
}

func (r *ShopRepo) Save(data interface{}) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("AssistantWife").C("Shop")

	model := data.(*models.ShopModel)

	if model.Id == "" {
		model.Id = bson.NewObjectId()
	}

	_, err = c.UpsertId(model.Id, model)
	if err != nil {
		panic(err)
	}
}

func (r ShopRepo) Get() (interface{}, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("AssistantWife").C("Shop")

	result := []models.ShopModel{}
	err = c.Find(bson.M{}).All(&result)

	return result, err
}

func (r ShopRepo) DeleteById(id interface{}) error {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("AssistantWife").C("Shop")

	idObj := id.(string)
	err = c.RemoveId(bson.ObjectIdHex(idObj))

	return err
}

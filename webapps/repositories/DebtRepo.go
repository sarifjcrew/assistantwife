package repositories

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"pribadi/assistantwife/webapps/models"
)

type DebtRepo struct {
}

func (r *DebtRepo) Save(data interface{}) error {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("AssistantWife").C("Debt")

	model := data.(*models.DebtModel)

	if model.Id == "" {
		model.Id = bson.NewObjectId()
	}

	_, err = c.UpsertId(model.Id, model)
	if err != nil {
		panic(err)
	}

	return err
}

func (r DebtRepo) Get() (interface{}, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("AssistantWife").C("Debt")

	result := []models.DebtModel{}
	err = c.Find(bson.M{}).All(&result)

	return result, err
}

func (r DebtRepo) DeleteById(id interface{}) error {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("AssistantWife").C("Debt")

	idObj := id.(string)
	err = c.RemoveId(bson.ObjectIdHex(idObj))

	return err
}

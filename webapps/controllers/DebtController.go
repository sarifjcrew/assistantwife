package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"pribadi/assistantwife/webapps/models"
)

type DebtController struct{}

func (c *DebtController) Save(w http.ResponseWritter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	defer r.Body.Close()
	br := bytes.NewReader(reqBody)
	decoder := json.NewDecoder(br)

	model := models.DebtModel{}
	_ = decoder.Decode(model)

}

package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"pribadi/assistantwife/webapps/models"
	"pribadi/assistantwife/webapps/repositories"
	"pribadi/assistantwife/webapps/services"
)

type SaldoController struct{}

func (c *SaldoController) Get(w http.ResponseWriter, r *http.Request) {
	repo := repositories.SaldoRepo{}
	service := services.SaldoService{}

	result, err := service.Get(repo)
	if err != nil {
		panic(err)
	}

	dataReturn, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(dataReturn)
}

func (c *SaldoController) Save(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	defer r.Body.Close()
	br := bytes.NewReader(reqBody)
	decoder := json.NewDecoder(br)

	model := new(models.SaldoModel)
	_ = decoder.Decode(model)

	repo := repositories.SaldoRepo{}
	service := services.SaldoService{}

	service.Save(model, repo)

	result, _ := json.Marshal(model)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}

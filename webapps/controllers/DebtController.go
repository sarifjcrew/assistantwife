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

type DebtController struct{}

func (c *DebtController) Get(w http.ResponseWriter, r *http.Request) {
	repo := repositories.DebtRepo{}
	service := services.DebtService{}

	result, err := service.Get(repo)
	if err != nil {
		panic(err)
	}

	dataReturn, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(dataReturn)
}

func (c *DebtController) Save(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	defer r.Body.Close()
	br := bytes.NewReader(reqBody)
	decoder := json.NewDecoder(br)

	model := new(models.DebtModel)
	_ = decoder.Decode(model)

	repo := repositories.DebtRepo{}
	service := services.DebtService{}

	service.Save(model, repo)

	result, _ := json.Marshal(model)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}

func (c *DebtController) Delete(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	defer r.Body.Close()
	br := bytes.NewReader(reqBody)
	decoder := json.NewDecoder(br)

	param := map[string]interface{}{}
	_ = decoder.Decode(&param)

	repo := repositories.DebtRepo{}
	service := services.DebtService{}

	err := service.DeleteById(repo, param["id"])

	result := Result{
		data: "",
		err:  err,
	}

	rensp, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(rensp)
}

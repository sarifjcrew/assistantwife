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

type Result struct {
	data interface{}
	err  error
}
type ShopController struct{}

func (c *ShopController) Save(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	defer r.Body.Close()
	br := bytes.NewReader(reqBody)
	decoder := json.NewDecoder(br)

	shopM := models.ShopModel{}
	_ = decoder.Decode(&shopM)

	repo := repositories.ShopRepo{}
	repo.Save(&shopM)

	result, _ := json.Marshal(shopM)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}

func (c *ShopController) Get(w http.ResponseWriter, r *http.Request) {
	repo := repositories.ShopRepo{}
	service := services.ShopService{}

	result, err := service.Get(repo)
	if err != nil {
		panic(err)
	}

	result1, _ := json.Marshal(result.([]models.ShopModel))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(result1)
}

func (c *ShopController) Delete(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	defer r.Body.Close()
	br := bytes.NewReader(reqBody)
	decoder := json.NewDecoder(br)

	param := map[string]interface{}{}

	_ = decoder.Decode(&param)

	repo := repositories.ShopRepo{}
	service := services.ShopService{}

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

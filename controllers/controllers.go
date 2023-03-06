package controllers

import (
	"ApiRestGolang/database"
	"ApiRestGolang/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersonalidades(w http.ResponseWriter, r *http.Request) {

	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func PersonalidadesPorId(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		panic(err.Error())
	}

	var p []models.Personalidade
	database.DB.First(&p, idNumber)

	json.NewEncoder(w).Encode(p)

}

func UpdatePersonalidade(w http.ResponseWriter, r *http.Request) {

	var personalidade models.Personalidade
	vars := mux.Vars(r)
	id := vars["id"]
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		panic(err.Error())
	}
	database.DB.First(&personalidade, idNumber)
	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Save(&personalidade)

	json.NewEncoder(w).Encode(&personalidade)
}

func RegisterPersonalidade(w http.ResponseWriter, r *http.Request) {

	var personalidadeNova models.Personalidade
	json.NewDecoder(r.Body).Decode(&personalidadeNova)
	database.DB.Create(&personalidadeNova)

	json.NewEncoder(w).Encode(personalidadeNova)
}

func DeletePersonalidade(w http.ResponseWriter, r *http.Request) {

	var personalidadeNova models.Personalidade
	vars := mux.Vars(r)
	id := vars["id"]
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		panic(err.Error())
	}
	var p []models.Personalidade
	database.DB.First(&p, idNumber)
	database.DB.Delete(&personalidadeNova, idNumber)
	json.NewEncoder(w).Encode(p)
}

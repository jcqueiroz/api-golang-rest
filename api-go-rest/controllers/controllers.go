package controllers

import (
	"ALURA/jcqueiroz2/api-go-rest/database"
	"ALURA/jcqueiroz2/api-go-rest/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personality
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)

}

func ReturnOnePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var Personality models.Personality
	database.DB.First(&Personality, id)
	json.NewEncoder(w).Encode(Personality)

}

func CreateNewPersonality(w http.ResponseWriter, r *http.Request) {
	var newPersonality models.Personality
	json.NewDecoder(r.Body).Decode(&newPersonality)
	database.DB.Create(&newPersonality)
	json.NewEncoder(w).Encode(newPersonality)
}

func DeleteOnePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var Personality models.Personality
	database.DB.Delete(&Personality, id)
	json.NewEncoder(w).Encode(Personality)
}

func EditPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var Personality models.Personality
	database.DB.First(&Personality, id)
	json.NewDecoder(r.Body).Decode(&Personality)
	database.DB.Save(&Personality)
	json.NewEncoder(w).Encode(Personality)

}

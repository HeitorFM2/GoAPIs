package resource

import (
	"api-mux/configs"
	"api-mux/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var users []models.Users
var user models.Users

func GetUsers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "applicattion/json")
	configs.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicattion/json")
	params := mux.Vars(r)
	configs.DB.First(&users, params["id"])
	// configs.DB.Where("email = ?", params["id"]).Find(&users)
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicattion/json")

	json.NewDecoder(r.Body).Decode(&users)
	result := configs.DB.Create(&user)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("InternalServerError"))
		return
	}

	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicattion/json")
	var user models.Users
	params := mux.Vars(r)

	configs.DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	result := configs.DB.Save(&user)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("InternalServerError"))
		return
	}

	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicattion/json")
	params := mux.Vars(r)
	result := configs.DB.Delete(&users, params["id"])
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("InternalServerError"))
		return
	}
	json.NewEncoder(w).Encode("The user " + params["id"] + "  has been deleted!")
}

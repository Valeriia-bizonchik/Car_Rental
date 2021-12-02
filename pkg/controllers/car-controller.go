package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Valeriia-bizonchik/CarRental/pkg/models"
	"github.com/Valeriia-bizonchik/CarRental/pkg/utils"
	"github.com/gorilla/mux"
)

var NewCar models.Car

func GetCar(w http.ResponseWriter, r *http.Request) {
	newCars := models.GetAllCars()
	res, _ := json.Marshal(newCars)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCarById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId := vars["carId"]
	ID, err := strconv.ParseInt(carId, 0, 0)
	if err != nil {
		fmt.Println("err while parsing")
	}
	carDetails, _ := models.GetCarById(ID)
	res, _ := json.Marshal(carDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateCar(w http.ResponseWriter, r *http.Request) {
	CreateCar := &models.Car{}
	utils.ParseBody(r, CreateCar)
	b := CreateCar.CreateCar()

	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId := vars["carId"]
	ID, err := strconv.ParseInt(carId, 0, 0)
	if err != nil {
		fmt.Println("err while parsing")
	}
	car := models.DeleteCar(ID)
	res, _ := json.Marshal(car)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	var updateCar = &models.Car{}
	utils.ParseBody(r, updateCar)
	vars := mux.Vars(r)
	carId := vars["carId"]
	ID, err := strconv.ParseInt(carId, 0, 0)
	if err != nil {
		fmt.Println("err while parsing")
	}
	carDetails, db := models.GetCarById(ID)

	if updateCar.Name != "" {
		carDetails.Name = updateCar.Name
	}

	db.Save(&carDetails)
	res, _ := json.Marshal(carDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

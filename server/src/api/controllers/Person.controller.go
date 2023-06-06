package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"pruebaTecnica/src/api/utils"
	"pruebaTecnica/src/core/services"

	"github.com/gorilla/mux"
)

func (server *Server) GetPersons(w http.ResponseWriter, r *http.Request) {

	person := services.PersonServices{}

	priorities, err := person.FindAllPersons(server.DB)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, priorities, "Persons was consulted succesfully")
}

func (server *Server) CreatePerson(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	person := services.PersonServices{}

	err = json.Unmarshal(body, &person)

	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	personCreated, err := person.SavePerson(server.DB)

	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, personCreated.ID))
	utils.JSON(w, http.StatusCreated, personCreated, "Person was created successfully")
}

func (server *Server) UpdatePerson(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	person := services.PersonServices{}

	err = json.Unmarshal(body, &person)

	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	personUpdated, err := person.UpdatePerson(server.DB)

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, personUpdated.ID))
	utils.JSON(w, http.StatusCreated, personUpdated, "Person was created successfully")
}

func (server *Server) DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idPerson, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	person := services.PersonServices{}

	deletedPerson, err := person.UnactivedPerson(server.DB, idPerson)

	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	utils.JSON(w, http.StatusOK, deletedPerson, "Person deleted successfully")
}

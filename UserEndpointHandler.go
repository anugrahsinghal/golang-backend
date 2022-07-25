package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func (apiCfg apiConfig) endpointUsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		apiCfg.handlerGetUser(w, r)
	case http.MethodPost:
		apiCfg.handleCreateUser(w, r)
	case http.MethodPut:
		apiCfg.handlerUpdateUser(w, r)
	case http.MethodDelete:
		apiCfg.handlerDeleteUser(w, r)
	default:
		respondWithError(w, 404, errors.New("method not supported"))
	}
}

func (apiCfg apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type UserDto struct {
		Age      int    `json:"age"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	userDto := UserDto{}
	err := decoder.Decode(&userDto)
	if err != nil {
		respondWithError(w, 404, err)
		return
	}
	user, err := apiCfg.dbClient.CreateUser(userDto.Email, userDto.Password, userDto.Name, userDto.Age)
	if err != nil {
		respondWithError(w, 404, err)
		return
	}
	respondWithJSON(w, http.StatusCreated, user)
}

func (apiCfg apiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request) {

	email := strings.TrimPrefix(r.URL.Path, "/users/")
	err := apiCfg.dbClient.DeleteUser(email)
	if err != nil {
		respondWithError(w, 404, err)
		return
	}
	respondWithJSON(w, http.StatusOK, struct{}{} /*struct{} part is struct definition, further {} are for creation*/)
}

func (apiCfg apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	type UserUpdateDto struct {
		Age      int    `json:"age"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	userUpdateDto := UserUpdateDto{}
	err := decoder.Decode(&userUpdateDto)
	if err != nil {
		respondWithError(w, 404, err)
		return
	}

	email := strings.TrimPrefix(r.URL.Path, "/users/")

	user, err := apiCfg.dbClient.UpdateUser(email, userUpdateDto.Password, userUpdateDto.Name, userUpdateDto.Age)
	if err != nil {
		respondWithError(w, 404, err)
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

func (apiCfg apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimPrefix(r.URL.Path, "/users/")

	user, err := apiCfg.dbClient.GetUser(email)
	if err != nil {
		respondWithError(w, 404, err)
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

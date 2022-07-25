package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func (apiCfg apiConfig) endpointPostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		apiCfg.handlerRetrievePost(w, r)
	case http.MethodPost:
		apiCfg.handlerCreatePost(w, r)
	case http.MethodDelete:
		apiCfg.handlerDeletePost(w, r)
	default:
		respondWithError(w, http.StatusNotFound, errors.New("method not supported"))
	}
}

func (apiCfg apiConfig) handlerCreatePost(w http.ResponseWriter, r *http.Request) {
	type PostDTO struct {
		UserEmail string `json:"userEmail"`
		Text      string `json:"text"`
	}
	decoder := json.NewDecoder(r.Body)
	postDto := PostDTO{}
	err := decoder.Decode(&postDto)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	post, err := apiCfg.dbClient.CreatePost(postDto.UserEmail, postDto.Text)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusCreated, post)
}

func (apiCfg apiConfig) handlerDeletePost(w http.ResponseWriter, r *http.Request) {
	uuid := strings.TrimPrefix(r.URL.Path, "/posts/")

	err := apiCfg.dbClient.DeletePost(uuid)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}

func (apiCfg apiConfig) handlerRetrievePost(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimPrefix(r.URL.Path, "/posts/")

	postArr, err := apiCfg.dbClient.GetPosts(email)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusOK, postArr)
}

package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		Valid string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140

	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}
	filtered := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}
	rawInput := strings.Split(params.Body, " ")
	cleanInput := []string{}
	for _, s := range rawInput {
		lowered := strings.ToLower(s)
		if _, ok := filtered[lowered]; ok {
			cleanInput = append(cleanInput, "****")
		} else {
			cleanInput = append(cleanInput, s)
		}
	}

	respondWithJSON(w, http.StatusOK, returnVals{
		Valid: strings.Join(cleanInput, " "),
	})
}

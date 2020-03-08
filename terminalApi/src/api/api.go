package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"terminal_api/postgres"
)

type UserData struct {
	Name		string	`json:"name"`
	LocationId	string	`json:"locationId"`
	TerminalId	string	`json:"terminalId"`
}

func RegisterTerm(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err.Error())
		}
		userData := new(UserData)
		err = json.Unmarshal(data, &userData)
		if err != nil {
			log.Fatal(err)
		}
		postgres.RegisterNewUser(userData.Name, userData.LocationId, userData.TerminalId)
		w.Write([]byte("OK"))
	}
}

func UnregisterTerm(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err.Error())
		}
		userData := new(UserData)
		err = json.Unmarshal(data, &userData)
		if err != nil {
			log.Fatal(err)
		}
		postgres.UnregisterUser(userData.TerminalId)
		w.Write([]byte("OK"))
	}
}
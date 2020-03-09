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

type TermData struct {
	TerminalId	string	`json:"terminalId"`
}

// Регистрация терминала
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

// Возврат терминала сотрудником
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

// Это функция возвращает весь журнал действий
func AllTermData(w http.ResponseWriter, r *http.Request) {

	jsonData, err := json.Marshal(postgres.GetJournalData())
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// Функция возвращает Историю одного терминала
func InfoAboutTerm(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err.Error())
		}
		userData := new(TermData)
		err = json.Unmarshal(data, &userData)
		if err != nil {
			log.Fatal(err)
		}

		jsonData, err := json.Marshal(postgres.TermHistory(userData.TerminalId))
		if err != nil {
			log.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

// Функция возвращает инфо Кто сейчас использует терминал
func WhoUseTerminal(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err.Error())
		}
		userData := new(TermData)
		err = json.Unmarshal(data, &userData)
		if err != nil {
			log.Fatal(err)
		}

		jsonData, err := json.Marshal(postgres.FindTerminal(userData.TerminalId))
		if err != nil {
			log.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}
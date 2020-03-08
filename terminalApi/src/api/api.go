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

//// Эта функция возвращает весь журнал действий, но она почему-то не хочет работать
//func AllTermData() (interface{}, error) {
//
//	var jsonData []byte
//	jsonData, err := json.Marshal(postgres.GetJournalData())
//	if err != nil {
//		log.Println(err)
//	}
//	return string(jsonData), nil
//}

//// Функция возвращает всю историю об одном терминале
//func InfoAboutTerm(w http.ResponseWriter, r *http.Request) (interface{}, error) {
//
//	if r.Method == "POST" {
//		data, err := ioutil.ReadAll(r.Body)
//		if err != nil {
//			log.Fatal(err.Error())
//		}
//		userData := new(UserData)
//		err = json.Unmarshal(data, &userData)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		jsonData := postgres.TermHistory(userData.TerminalId)
//
//		return json.Marshal(jsonData)
//	}
//}

//// Функция возвращает инфо о том кто его использует в настоящее время
//func WhoUseTerminal(w http.ResponseWriter, r *http.Request) (interface{}, error) {
//
//	if r.Method == "POST" {
//		data, err := ioutil.ReadAll(r.Body)
//		if err != nil {
//			log.Fatal(err.Error())
//		}
//		userData := new(UserData)
//		err = json.Unmarshal(data, &userData)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		jsonData := postgres.FindTerminal(userData.TerminalId)
//
//		return json.Marshal(jsonData)
//	}
//}
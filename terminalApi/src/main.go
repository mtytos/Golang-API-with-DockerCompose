package main

import (
	"bytes"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"
	"terminal_api/api"
	"terminal_api/postgres"
)

// Проверка работоспособности метода по Регистрации выдачи ТСД
func test(w http.ResponseWriter, r *http.Request) {

	log.Println("Got test request")

	data := []byte(`{"name":"John","locationId":"911","terminalId":"109"}`)
	request := bytes.NewReader(data)
	_, err := http.Post("http://localhost:8080/registerTerm", "application/json", request)
	if err != nil {
		log.Fatal("Error: ", err)
	} else {
		w.Write([]byte("Сотрудник зарегистрирован в журнале, ТСД выдан!"))
		log.Println("Success")
	}
}

// Проерка работоспособности метода по регистрации Возврата ТСД
func test2(w http.ResponseWriter, r *http.Request) {

	log.Println("Got test request2")

	data := []byte(`{"name":"John","locationId":"911","terminalId":"109"}`)
	request := bytes.NewReader(data)
	_, err := http.Post("http://localhost:8080/UnregisterTerm", "application/json", request)
	if err != nil {
		log.Fatal("Error: ", err)
	} else {
		w.Write([]byte("Сотрудник сдал ТСД, данные в журнале обновлены!"))
		log.Println("Success2")
	}
}

// Проерка работоспособности метода
func test3(w http.ResponseWriter, r *http.Request) {

	log.Println("Got test request3")

	data := []byte(`{"terminalId":"109"}`)
	request := bytes.NewReader(data)
	_, err := http.Post("http://localhost:8080/InfoAboutTerm", "application/json", request)
	if err != nil {
		log.Fatal("Error: ", err)
	} else {
		w.Write([]byte("Журнал с историей терминала!"))
		log.Println("Success3")
	}
}


// Проерка работоспособности метода по регистрации Возврата ТСД
func test4(w http.ResponseWriter, r *http.Request) {

	response, err := api.AllTermData2()
	if err != nil {
		log.Fatal("Error: ", err)
	}
	fmt.Fprintf(w, string(response))
}


func main()  {
	postgres.InitDBTables()

	port := os.Getenv("API_PORT")

	router := chi.NewRouter()

	router.Route("/registerTerm", func(r chi.Router) {
		r.Post("/", api.RegisterTerm)
	})

	router.Route("/UnregisterTerm", func(r chi.Router) {
		r.Post("/", api.UnregisterTerm)
	})

	router.Route("/AllTermData", func(r chi.Router) {
		r.Post("/", api.AllTermData)
	})

	router.Route("/InfoAboutTerm", func(r chi.Router) {
		r.Post("/", api.InfoAboutTerm)
	})

	router.Route("/WhoUseTerminal", func(r chi.Router) {
		r.Post("/", api.WhoUseTerminal)
	})

	// Роуты для проверки
	router.Route("/test", func(r chi.Router) {
		r.Get("/", test)
	})

	router.Route("/test2", func(r chi.Router) {
		r.Get("/", test2)
	})

	router.Route("/test3", func(r chi.Router) {
		r.Get("/", test3)
	})

	router.Route("/test4", func(r chi.Router) {
		r.Get("/", test4)
	})

	http.ListenAndServe(":" + port, router)
}

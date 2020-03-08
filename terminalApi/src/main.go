package main

import (
	"bytes"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"
	"terminal_api/api"
	"terminal_api/postgres"
)

func test(w http.ResponseWriter, r *http.Request) {

	log.Println("Got test request")

	data := []byte(`{"name":"Pe","locationId":"969","terminalId":"676"}`)
	request := bytes.NewReader(data)
	_, err := http.Post("http://localhost:8080/registerTerm", "application/json", request)
	if err != nil {
		log.Fatal("Error: ", err)
	} else {
		w.Write([]byte("11 I will try to do something!"))
		log.Println("Success")
	}
}

func test2(w http.ResponseWriter, r *http.Request) {

	log.Println("Got test request2")

	data := []byte(`{"name":"Pe","locationId":"969","terminalId":"676"}`)
	request := bytes.NewReader(data)
	_, err := http.Post("http://localhost:8080/UnregisterTerm", "application/json", request)
	if err != nil {
		log.Fatal("Error: ", err)
	} else {
		w.Write([]byte("отправил данные для обновления!"))
		log.Println("Success2")
	}
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



	router.Route("/test", func(r chi.Router) {
		r.Get("/", test)
	})

	router.Route("/test2", func(r chi.Router) {
		r.Get("/", test2)
	})
	http.ListenAndServe(":" + port, router)
}

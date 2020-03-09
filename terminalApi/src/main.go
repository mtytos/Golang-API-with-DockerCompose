package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"os"
	"terminal_api/api"
	"terminal_api/postgres"
)


// ЧАСТЬ КОДА, КОТОРАЯ НУЖНА ДЛЯ ОТЛАДКИ И ДЕБАГА
//
//// Проверка работоспособности метода по Регистрации выдачи ТСД
//func test1(w http.ResponseWriter, r *http.Request) {
//
//	log.Println("Got test request")
//
//	data := []byte(`{"name":"ilya","locationId":"777","terminalId":"777"}`)
//	request := bytes.NewReader(data)
//	_, err := http.Post("http://localhost:8080/registerTerm", "application/json", request)
//	if err != nil {
//		log.Fatal("Error: ", err)
//	} else {
//		w.Write([]byte("Сотрудник зарегистрирован в журнале, ТСД выдан!"))
//		log.Println("Success")
//	}
//}
//
//// Проерка работоспособности метода по регистрации Возврата ТСД
//func test2(w http.ResponseWriter, r *http.Request) {
//
//	log.Println("Got test request2")
//
//	data := []byte(`{"name":"ilya","locationId":"777","terminalId":"777"}`)
//	request := bytes.NewReader(data)
//	_, err := http.Post("http://localhost:8080/UnregisterTerm", "application/json", request)
//	if err != nil {
//		log.Fatal("Error: ", err)
//	} else {
//		w.Write([]byte("Сотрудник сдал ТСД, данные в журнале обновлены!"))
//		log.Println("Success2")
//	}
//}
//
//// Проерка работоспособности метода "История одного терминала"
//func test3(w http.ResponseWriter, r *http.Request) {
//
//	var buffer []byte
//	buffer = make([]byte, 200, 200)
//
//	log.Println("Got test request3")
//
//	data := []byte(`{"terminalId":"777"}`)
//	request := bytes.NewReader(data)
//	response, err := http.Post("http://localhost:8080/InfoAboutTerm", "application/json", request)
//	if err != nil {
//		log.Fatal("Error: ", err)
//	} else {
//		response.Body.Read(buffer)
//		w.Write(buffer)
//		log.Println("Success3")
//	}
//}
//
//// Who using
//func test4(w http.ResponseWriter, r *http.Request) {
//
//	var buffer []byte
//	buffer = make([]byte, 200, 200)
//
//	log.Println("Got test request4")
//
//	data := []byte(`{"terminalId":"777"}`)
//	request := bytes.NewReader(data)
//	response, err := http.Post("http://localhost:8080/WhoUseTerminal", "application/json", request)
//	if err != nil {
//		log.Fatal("Error: ", err)
//	} else {
//		response.Body.Read(buffer)
//		w.Write(buffer)
//		log.Println("Success4")
//	}
//}
//
//// All journal
//func test5(w http.ResponseWriter, r *http.Request) {
//
//	var buffer []byte
//	buffer = make([]byte, 500, 500)
//	log.Println("Got test request5")
//	data := []byte(`{""}`)
//	request := bytes.NewReader(data)
//	response, err := http.Post("http://localhost:8080/AllTermData", "application/json", request)
//	if err != nil {
//		log.Fatal("Error: ", err)
//	} else {
//		response.Body.Read(buffer)
//		w.Write(buffer)
//		log.Println("Success5")
//	}
//}
//
//// КОНЕЦ

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

	// !!!!   РОУТЫ ДЛЯ ПРОВЕРКИ   !!!!!

	//router.Route("/test1", func(r chi.Router) {
	//	r.Get("/", test1)
	//})
	//
	//router.Route("/test2", func(r chi.Router) {
	//	r.Get("/", test2)
	//})
	//
	//router.Route("/test3", func(r chi.Router) {
	//	r.Get("/", test3)
	//})
	//
	//router.Route("/test4", func(r chi.Router) {
	//	r.Get("/", test4)
	//})
	//
	//router.Route("/test5", func(r chi.Router) {
	//	r.Get("/", test5)
	//})

	// КОНЕЦ

	http.ListenAndServe(":" + port, router)
}

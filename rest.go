package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"DMS_v1.1/utils"
	"github.com/gorilla/mux"
)

var port string

type url string

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Context-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func documentation(w http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
	}
	// rw.Header().Add("Context-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func getBarcode(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// case "GET":
	// 	blockchain.Status(blockchain.Blockchain(), w)
	case "POST":
		var reqMsg reqMsgGenBarcode
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&reqMsg))
		respMsg := GenerateBarcode(reqMsg)
		fmt.Println(respMsg)
		w.WriteHeader(http.StatusCreated) // 201 : created
	}
}

func getBarcodeUsingID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var reqMsg reqMsgSearchBarcode
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&reqMsg))
		respMsg := SearchBarcode(reqMsg)
		fmt.Println(respMsg)
		w.WriteHeader(http.StatusCreated) // 201 : created
	}
}

func getMypage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var reqMsg reqMsgGetMypage
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&reqMsg))
		respMsg := BarcodeRouter(reqMsg.Barcode)
		fmt.Println(respMsg)
		w.WriteHeader(http.StatusCreated) // 201 : created
	}
}

func getBarcodeImg(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var reqMsg reqMsgSearchBarcode
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&reqMsg))
		GenerateBarcodeImg(reqMsg.ID)
		w.WriteHeader(http.StatusCreated) // 201 : created
	}
}

func RestStart(aPort int) {
	port = fmt.Sprintf(":%d", aPort)
	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddleware, loggerMiddleware)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/barcode", getBarcode).Methods("POST")
	router.HandleFunc("/search", getBarcodeUsingID).Methods("POST")
	router.HandleFunc("/mypage", getMypage).Methods("POST")

	router.HandleFunc("/genimg", getBarcodeImg).Methods("POST")

	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

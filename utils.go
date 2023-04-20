package main

import (
	"encoding/json"
	"log"
)

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func ToJson(i interface{}) []byte {
	r, err := json.Marshal(i)
	HandleErr(err)
	return r
}

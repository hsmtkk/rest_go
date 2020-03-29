package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const URL = "http://dummy.restapiexample.com/api/v1/create"

type request struct {
	Name   string `json:"name"`
	Salary int    `json:"salary"`
	Age    int    `json:"age"`
}

type response struct {
	Status string `json:"status"`
	Data   data   `json:"data"`
}

type data struct {
	Name   string `json:"name"`
	Salary int    `json:"salary"`
	Age    int    `json:"age"`
	ID     int    `json:"id"`
}

func main() {
	req := request{Name: "test", Salary: 123, Age: 23}
	reqBytes, err := json.MarshalIndent(&req, " ", " ")
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post(URL, "application/json", bytes.NewReader(reqBytes))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(respBytes))
	res := response{}
	if err := json.Unmarshal(respBytes, &res); err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

package create

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const createURL = "http://dummy.restapiexample.com/api/v1/create"

type Request struct {
	Name   string `json:"name"`
	Salary int    `json:"salary"`
	Age    int    `json:"age"`
}

type Response struct {
	Status string `json:"status"`
	Data   data   `json:"data"`
}

type data struct {
	Name   string `json:"name"`
	Salary int    `json:"salary"`
	Age    int    `json:"age"`
	ID     int    `json:"id"`
}

type Creator interface {
	Create(Request) (Response, string, string, error)
}

func New() Creator {
	return NewWithClient(http.DefaultClient, createURL)
}

func NewWithClient(client *http.Client, url string) Creator {
	return &creatorImpl{client, url}
}

type creatorImpl struct {
	client *http.Client
	url    string
}

func (c *creatorImpl) Create(req Request) (Response, string, string, error) {
	reqBytes, err := json.Marshal(&req)
	if err != nil {
		return Response{}, "", "", err
	}
	resp, err := http.Post(c.url, "application/json", bytes.NewReader(reqBytes))
	if err != nil {
		return Response{}, string(reqBytes), "", err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{}, string(reqBytes), "", err
	}
	res := Response{}
	if err := json.Unmarshal(respBytes, &res); err != nil {
		return Response{}, string(reqBytes), string(respBytes), err
	}
	return res, string(reqBytes), string(respBytes), nil
}

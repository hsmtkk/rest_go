package create_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hsmtkk/rest_go/pkg/create"
	"github.com/stretchr/testify/assert"
)

func TestReal(t *testing.T) {
	c := create.New()
	req := create.Request{Name: "alpha", Salary: 123, Age: 10}
	res, reqBody, resBody, err := c.Create(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	fmt.Println(reqBody)
	fmt.Println(resBody)
}

const wantResponse = `{"status":"success","data":{"name":"alpha","salary":123,"age":10,"id":70}}`

func TestStub(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, wantResponse)
	}))
	defer ts.Close()

	c := create.NewWithClient(ts.Client(), ts.URL)
	req := create.Request{Name: "alpha", Salary: 123, Age: 10}
	got, _, _, err := c.Create(req)
	assert.Nil(t, err, "error should not occur")
	assert.Equal(t, got.Status, "success", "should match")
	assert.Equal(t, got.Data.Name, "alpha", "should match")
	assert.Equal(t, got.Data.Salary, 123, "should match")
	assert.Equal(t, got.Data.Age, 10, "should match")
	assert.Equal(t, got.Data.ID, 70, "should match")
}

package tests

import (
	"testing"
	"net/http"
	"log"
	"io/ioutil"
	fmt "fmt"
)

func TestListCart(t *testing.T) {
	id,TokenLogin := Login("prayugo@gmail.com", "prayugo")
	if id == "" {
		t.Error("Login failed")
	}
	url := "http://localhost:8000/listcart"
	queryUrl := url + "?id=" + id
	Client := &http.Client{}
	fmt.Println(string(TokenLogin))
	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		log.Panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", TokenLogin)
	resp, err := Client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	Body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	t.Log(string(Body))
	
}




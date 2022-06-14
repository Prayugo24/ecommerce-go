package tests

import (
	"testing"
	"net/http"
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
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", TokenLogin)
	resp, err := Client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	Body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(Body))
	
}




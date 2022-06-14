package tests

import (
	"testing"
	"net/http"
	"log"
	"io/ioutil"
)
func TestSearch(t *testing.T) {
	url := "http://localhost:8000/users/search?name=Alienware"
	Client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
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
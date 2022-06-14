package tests

import (
	"testing"
	"net/http"
	"log"
	"io/ioutil"
	"strings"
)
func TestEditHomeAddress(t *testing.T) {
	id,TokenLogin := Login("prayugo@gmail.com", "prayugo")
	if id == "" {
		t.Error("Login failed")
	}
	url := "http://localhost:8000/edithomeaddress"
	queryUrl := url + "?id=" + id
	rawDataJson := `{
		"house_name": "Rumah Bahagia",
		"street": "Jalan Tirto Sari",
		"city_name": "Purwokerto",
		"pin_code": "20112"
	  }`
	Client := &http.Client{}
	// PUT
	req, err := http.NewRequest("PUT", queryUrl, strings.NewReader(rawDataJson))
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
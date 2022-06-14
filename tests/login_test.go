package tests

import (
	"testing"
	"net/http"
	"strings"
	"io/ioutil"
	"encoding/json"
	
)

func Login(email string, password string) (string,string) {
	url := "http://localhost:8000/users/login"

	rawDataJson := `{"email":"` + email + `","password":"` + password + `"}`

	Client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(rawDataJson))
	if err != nil {
		return "", ""
	}
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := Client.Do(req)
	if err != nil {
		return "", ""
	}
	defer resp.Body.Close()
	Body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", ""
	}
	var data map[string]interface{}
	err = json.Unmarshal(Body, &data)
	if err != nil {
		return "", ""
	}
	id := data["_id"].(string)
	token := data["refresh_token"].(string)
	return id, token

}

func TestLogin(t *testing.T) {
	_,token := Login("prayugo@gmail.com", "prayugo")
	if token == "" {
		t.Error("Login failed")
	}
	t.Log(token)
}
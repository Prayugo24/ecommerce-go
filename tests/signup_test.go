package tests
import (
	"testing"
	"net/http"
	"log"
	"io/ioutil"
	"strings"
	
)

func TestSignUp(t *testing.T) {
	url := "http://localhost:8000/users/signup"
	rawDataJson := `{
			"first_name": "Prayugo",
			"last_name": "Dwi",
			"email": "prayugo@gmail.com",
			"password": "prayugo",
			"phone": "+08965456789"
	}`
	Client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(rawDataJson))
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
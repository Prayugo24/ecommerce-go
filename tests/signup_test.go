package tests
import (
	"testing"
	"net/http"
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
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
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
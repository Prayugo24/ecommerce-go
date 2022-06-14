package tests

import (
	"testing"
	"net/http"
	"io/ioutil"
	"strings"
)
func TestAddProduct(t *testing.T) {
	url := "http://localhost:8000/admin/addproduct"
	rawDataJson := `{
		"product_name": "Alienware x10",
		"price": 2500,
		"rating": 10,
		"image": "alienware.jpg"
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
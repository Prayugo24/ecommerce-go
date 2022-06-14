package tests

import (
	"testing"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func listProduct() string{
	url := "http://localhost:8000/users/productview"
	Client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ""
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := Client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	Body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	if resp.StatusCode != 200 {
		return ""
	}
	var data []map[string]interface{}
	err = json.Unmarshal(Body, &data)
	if err != nil {
		return ""
	}
	id := data[0]["Product_ID"].(string)
	return id
}

func TestListProduct(t *testing.T) {
	id := listProduct()
	if id == "" {
		t.Error("Failed to list product")
	}
	t.Log(id)
}
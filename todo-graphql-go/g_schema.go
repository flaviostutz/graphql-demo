package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
)

func buildTodoSchema() (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    buildQuery(),
		Mutation: buildMutation(),
	})
}

func sendJSONRequest(method string, url string, input map[string]interface{}) (*http.Response, error) {
	var di = bytes.NewBuffer([]byte(""))
	fmt.Println("HEREaaaa")
	if input != nil {
		data, err := json.Marshal(input)
		if err != nil {
			return nil, err
		}
		di = bytes.NewBuffer(data)
	}

	fmt.Println("HEREbbbb")
	req, err := http.NewRequest(method, url, di)
	if err != nil {
		return nil, err
	}
	fmt.Println("HEREccc")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	return client.Do(req)
}

func processResponse(response *http.Response, validStatus int, array bool) (interface{}, error) {
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Couldn't process response. err=%s", err)
	}

	if response.StatusCode != validStatus {
		return nil, fmt.Errorf("Couldn't process response. status=%d. expected=%d", response.StatusCode, validStatus)
	}

	var resdata interface{}
	resdata = make(map[string]interface{})
	if array {
		resdata = make([]map[string]interface{}, 0)
	}
	err = json.Unmarshal(data, &resdata)
	if err != nil {
		return nil, fmt.Errorf("Couldn't process response. err=%s", err)
	}

	return resdata, nil
}

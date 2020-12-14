package graph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/flaviostutz/graphql-demo/todo-graphql-gqlgen/graph/model"
)

func sendJSONRequest(method string, url string, input interface{}) (*http.Response, error) {
	var di = bytes.NewBuffer([]byte(""))
	if input != nil {
		data, err := json.Marshal(input)
		if err != nil {
			return nil, err
		}
		di = bytes.NewBuffer(data)
	}

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

func processResponseTodo(response *http.Response, validStatus int) (*model.Todo, error) {
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Couldn't process response. err=%s", err)
	}

	if response.StatusCode != validStatus {
		return nil, fmt.Errorf("Couldn't process response. status=%d. expected=%d", response.StatusCode, validStatus)
	}

	result := model.Todo{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, fmt.Errorf("Couldn't process response. err=%s", err)
	}

	return &result, nil
}

func processResponseList(response *http.Response, validStatus int) ([]*model.Todo, error) {
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Couldn't process response. err=%s", err)
	}

	if response.StatusCode != validStatus {
		return nil, fmt.Errorf("Couldn't process response. status=%d. expected=%d", response.StatusCode, validStatus)
	}

	result := []*model.Todo{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, fmt.Errorf("Couldn't process response. err=%s", err)
	}

	return result, nil
}

func processResponseText(response *http.Response, validStatus int) (string, error) {
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("Couldn't process response. err=%s", err)
	}

	if response.StatusCode != validStatus {
		return "", fmt.Errorf("Couldn't process response. status=%d. expected=%d", response.StatusCode, validStatus)
	}

	return string(data), nil
}

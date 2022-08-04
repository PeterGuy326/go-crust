package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	CONTENT_TYPE string        = "application/json; charset=utf-8"
	TIME_OUT     time.Duration = 5 * time.Second
)

const (
	METHOD_POST = "POST"
)

type API struct {
	Host   string
	App    string
	Action string
	Type   string
}

func Invoke(api API, function string, params interface{}) (string, error) {
	url := fmt.Sprintf("%s%s%s%s", api.Host, api.App, api.Action, function)
	jsonStr, _ := json.Marshal(params)
	req, err := http.NewRequest(METHOD_POST, url, bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", CONTENT_TYPE)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: TIME_OUT}
	resp, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)

	return string(result), nil
}

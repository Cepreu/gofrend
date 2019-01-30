package main

// Inspired by http://www.shaneword.com/post/2015/12/15/go-simple-soap-web-service-client.html

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Cepreu/gofrend/f9configurator"
)

var conf = struct {
	F9URL           string `json:"f9url"`
	F9Authorization string `json:"f9credentials"`
}{}

func init() {
	var err error

	err = f9configurator.GetConfiguration(f9configurator.JSON, &conf, "config.json")
	if err != nil {
		return
	}
}

func queryF9(generateRequestContent func() string) ([]byte, error) {
	url := conf.F9URL
	client := &http.Client{}
	sRequestContent := generateRequestContent()

	requestContent := []byte(sRequestContent)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestContent))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("Accept", "text/xml")
	req.Header.Add("Authorization", "Basic "+conf.F9Authorization)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(req)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("Error Respose " + resp.Status)
	}
	contents, err := ioutil.ReadAll(resp.Body)
	return contents, err
}

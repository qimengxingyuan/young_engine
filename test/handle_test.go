package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

var httpCli *http.Client

func init() {
	httpCli = &http.Client{}
}

func getResp(resp *http.Response, target interface{}) error {
	if resp == nil {
		err := errors.New("empty response")
		return err
	} else if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		err := fmt.Errorf("bad response. status_code=%d, msg=%s", resp.StatusCode, string(body))
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, target)
	return err
}

type Resp struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func TestPing(t *testing.T) {
	url := "http://127.0.0.1:8888/ping"
	resp, err := httpCli.Get(url)
	if err != nil {
		t.Error(err)
	} else {
		response := Resp{}
		err = getResp(resp, &response)
		if err != nil {
			t.Error(err)
		} else {
			t.Log(response.Data)
		}
	}
}

package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type SlackApiClient struct {
	httpClient *http.Client
}
type ApiClient interface {
	PostCustom(url string, payload []byte, token string) ([]byte, error) 
}

func NewSlackApiClient() ApiClient{
	return &SlackApiClient{&http.Client{
		Timeout: 60 * time.Second,
	},
  }
}

func (sc *SlackApiClient) PostCustom(url string, payload []byte, token string) ([]byte, error) {
    // convert byte slice to io.Reader
    reader := bytes.NewReader(payload)
	req, err := http.NewRequest("POST",url, reader)
	if err != nil {
		return nil, err
	}
	
	req.Header.Add("Content-Type", `application/json`)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := sc.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	return body, nil
}

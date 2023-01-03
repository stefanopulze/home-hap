package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"home-hap/pkg/logging"
	"io"
	"net/http"
)

type HttpRestClient struct {
	endpoint string
	log      *logging.Logger
}

func NewRestClient(endpoint string) *HttpRestClient {
	return &HttpRestClient{
		endpoint: endpoint,
		log:      logging.GetLoggerWithField("service", "rest-client"),
	}
}

func (c *HttpRestClient) Get(path string, body io.Reader) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.endpoint, path)
	c.log.Debug(fmt.Sprintf("Calling: %s", url))

	request, err := http.NewRequest(http.MethodGet, url, body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	c.log.Debug(fmt.Sprintf("Response: %s", string(b)))
	return b, nil
}

func (c *HttpRestClient) PostJson(path string, data interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.endpoint, path)
	c.log.Debug(fmt.Sprintf("Calling: %s", url))

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	c.log.Debug(fmt.Sprintf("Response: %s", string(b)))

	return b, nil
}

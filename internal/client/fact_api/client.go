package fact_api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	BaseURL = "https://development.kpi-drive.ru/_api"
)

type Client struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

func NewClient(APIKey string) *Client {
	return &Client{
		BaseURL: BaseURL,
		APIKey:  APIKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) sendRequest(req *http.Request) error {
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			// TODO: Handle error
		}
	}(res.Body)

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	return nil
}

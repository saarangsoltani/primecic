package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gocandoit.xyz/model"
)

type APIClient struct {
	BaseURL string
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{BaseURL: baseURL}
}

func (c *APIClient) SendRequest(method, path string, payload interface{}) ([]byte, error) {
	var token string = os.Getenv("API_TOKEN")
	var requestBody []byte
	if payload != nil {
		var err error
		requestBody, err = json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("error marshaling JSON payload: %v", err)
		}
	}

	req, err := http.NewRequest(method, c.BaseURL+path, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected response status code: %d", resp.StatusCode)
	}

	return []byte(body), nil
}

func GetApplicant(applicantID string) (*model.Applicant, error) {
	apiHost := os.Getenv("API_HOST")
	client := NewAPIClient(apiHost)
	response, err := client.SendRequest("GET", "/applicants/"+applicantID, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var strapiResp model.StrapiResp
	json.Unmarshal(response, &strapiResp)
	return &strapiResp.Data.Attributes, nil

}

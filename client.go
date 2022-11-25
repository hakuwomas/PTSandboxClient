package services

import (
	"PTSandboxClient/pkg/models"
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	host       string
	httpClient *http.Client
	apiKey     string
}

func NewClient(host string, apiKey string, timeout time.Duration) *Client {
	client := &http.Client{
		Timeout: timeout,
	}

	return &Client{
		host:       host,
		httpClient: client,
		apiKey:     apiKey,
	}
}

func (c *Client) UploadScanFile(ctx context.Context, filepath string) (*models.Response, error) {

	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.host+"/storage/uploadScanFile", bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	res := models.Response{}
	if err := c.sendRequest(req, &res, "application/octet-stream"); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) CreateScanTask(ctx context.Context, options *models.CreateScanTaskOptions) (*models.Response, error) {

	data, _ := json.Marshal(options)

	req, err := http.NewRequest("POST", c.host+"/analysis/createScanTask", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	res := models.Response{}
	if err := c.sendRequest(req, &res, "application/json"); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetReport(ctx context.Context, taskId string) (*models.Response, error) {

	req, err := http.NewRequest("POST", c.host+"/analysis/report", bytes.NewBuffer([]byte(`{
																									"scan_id": "`+taskId+`"
																								}`)))

	if err != nil {
		return nil, err
	}

	res := models.Response{}
	if err := c.sendRequest(req, &res, "application/json"); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) CheckTask(ctx context.Context, taskId string) (*models.Response, error) {

	req, err := http.NewRequest("POST", c.host+"/analysis/checkTask", bytes.NewBuffer([]byte(`{
																									"scan_id": "`+taskId+`"
																								}`)))

	if err != nil {
		return nil, err
	}

	res := models.Response{}
	if err := c.sendRequest(req, &res, "application/json"); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) sendRequest(req *http.Request, v interface{}, contentType string) error {
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Accept", "application/vnd.ptsecurity.app-v2")
	req.Header.Set("X-API-Key", c.apiKey)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	response := models.Response{}
	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return err
	}

	return nil
}

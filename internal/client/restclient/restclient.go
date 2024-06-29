package restclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/KozlovNikolai/CMDorders/internal/models"
)

type RestClient struct {
	BaseURL    string
	HTTPClient *http.Client
	Target     string
	Model      interface{}
}

func NewRestClient(baseURL string, target string, model interface{}) RestClient {
	return RestClient{
		HTTPClient: &http.Client{},
		BaseURL:    baseURL,
		Target:     target,
		Model:      model,
	}
}

func (c RestClient) GetByID(ctx context.Context, id uint64) (interface{}, error) {
	url := fmt.Sprintf("%s%s%d", c.BaseURL, c.Target, id)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get response, status code: %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	m, err := processValue(c.Model)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c RestClient) GetList(ctx context.Context) (interface{}, error) {
	url := fmt.Sprintf("%s%slist", c.BaseURL, c.Target)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get response, status code: %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)

	m, err := processValue(c.Model)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func processValue(value interface{}) (interface{}, error) {
	switch v := value.(type) {
	case models.Patient:
		return v, nil
	case models.Service:
		return v, nil
	default:
		fmt.Printf("Unknown type: %T\n", v)
		return nil, fmt.Errorf("unknown structure")
	}
}

// func (c *RestClient) CreateEmployer(employer modelss.Employer) (*modelss.Employer, error) {
// 	url := fmt.Sprintf("%s/employers", c.BaseURL)
// 	jsonValue, _ := json.Marshal(employer)
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
// 	if err != nil {
// 		return nil, err
// 	}
// 	req.Header.Set("Content-Type", "application/json")

// 	resp, err := c.RestClient.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusCreated {
// 		return nil, fmt.Errorf("failed to create employer, status code: %d", resp.StatusCode)
// 	}

// 	body, _ := ioutil.ReadAll(resp.Body)
// 	var createdEmployer modelss.Employer
// 	err = json.Unmarshal(body, &createdEmployer)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &createdEmployer, nil
// }

// func (c *RestClient) UpdateEmployer(id int, employer modelss.Employer) error {
// 	url := fmt.Sprintf("%s/employers/%d", c.BaseURL, id)
// 	jsonValue, _ := json.Marshal(employer)
// 	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
// 	if err != nil {
// 		return err
// 	}
// 	req.Header.Set("Content-Type", "application/json")

// 	resp, err := c.RestClient.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return fmt.Errorf("failed to update employer, status code: %d", resp.StatusCode)
// 	}

// 	return nil
// }

// func (c *RestClient) DeleteEmployer(id int) error {
// 	url := fmt.Sprintf("%s/employers/%d", c.BaseURL, id)
// 	req, err := http.NewRequest("DELETE", url, nil)
// 	if err != nil {
// 		return err
// 	}

// 	resp, err := c.RestClient.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return fmt.Errorf("failed to delete employer, status code: %d", resp.StatusCode)
// 	}

// 	return nil
// }

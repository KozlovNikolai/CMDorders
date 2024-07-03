package restclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/KozlovNikolai/CMDorders/internal/models"
	"go.uber.org/zap"
)

type RestClient struct {
	Logger     *zap.Logger
	BaseURL    string
	HTTPClient *http.Client
	Target     string
	Model      interface{}
}

func NewRestClient(baseURL string, target string, model interface{}, logger *zap.Logger) RestClient {
	return RestClient{
		Logger:     logger,
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
	m, err := processValue(c.Model, c.Logger)
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

	m, err := processValue(c.Model, c.Logger)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func processValue(value interface{}, logger *zap.Logger) (interface{}, error) {
	switch v := value.(type) {
	case models.Patient:
		return v, nil
	case models.Service:
		return v, nil
	default:
		msg := fmt.Sprintf("Unknown type: %T\n", v)
		logger.Debug("processValue",
			zap.String("info", msg),
		)
		return nil, fmt.Errorf("unknown structure")
	}
}

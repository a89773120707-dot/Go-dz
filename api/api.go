package api

import (
	"3-bin_manager/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	httpClient *http.Client   // стандартный http клиент
	cfg        *config.Config // структура cfg
	baseUrl    string         // url api.jsonbin
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
		httpClient: &http.Client{},                // экземпляр
		cfg:        cfg,                           // apiKey внутри стурктуры метод apiKey воспользоваться
		baseUrl:    "https://api.jsonbin.io/v3/b", // baseurl
	}
}

func (c *Client) CreateBin(data map[string]interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("ошибка сериализации JSON: %w", err)
	}

	req, err := http.NewRequest("POST", c.baseUrl, bytes.NewReader(jsonData))
	if err != nil {
		return "", fmt.Errorf("создание запроса: %w", err)
	}

	c.setHeders(req, true)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("соединение: %w", err)
	}
	defer resp.Body.Close()
	if err := c.checkResponse(resp); err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения ответа: %w", err)
	}
	var result struct {
		Metadata struct {
			ID string `json:"id"`
		} `json:"metadata"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("парсинг ответа: %w", err)
	}
	if result.Metadata.ID == "" {
		return "", fmt.Errorf("сервер не вернул ID бина")
	}
	return result.Metadata.ID, nil
}

func (c *Client) GetBin(binID string) (map[string]interface{}, error) {
	URL := c.baseUrl + "/" + binID
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания запроса: %w", err)
	}
	c.setHeders(req, false)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("соединение: %w", err)
	}
	defer resp.Body.Close()

	if err := c.checkResponse(resp); err != nil {
		return nil, fmt.Errorf("ошибка соединения: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения ответа: %w", err)
	}
	var result struct {
		Record map[string]interface{} `json:"record"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("ошибка обработки ответа сервера: %w", err)
	}
	if result.Record == nil {
		return nil, fmt.Errorf("Бин Пуст")
	}

	return result.Record, nil

}

func (c *Client) UpdateBin(binID string, data map[string]interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("Ошибка конвертации в json: %w", err)
	}
	URL := c.baseUrl + "/" + binID

	req, err := http.NewRequest(http.MethodPut, URL, bytes.NewReader(jsonData))
	if err != nil {
		return "", fmt.Errorf("ошибка создания запроса: %w", err)
	}

	c.setHeders(req, true)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("соединение: %w", err)
	}
	defer resp.Body.Close()

	if err := c.checkResponse(resp); err != nil {
		return "", fmt.Errorf("ошибка соединения: %w", err)
	}

	var result struct {
		Metadata struct {
			ID string `json:"id"`
		} `json:"metadata"`
	}

	dec := json.NewDecoder(resp.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&result); err != nil {
		return "", fmt.Errorf("Ошибка декодирования: %w", err)
	}

	if result.Metadata.ID == "" {
		return "", fmt.Errorf("сервер не вернул ID бина")
	}
	return result.Metadata.ID, nil
}
func (c *Client) DeleteBin(binID string) error {
	if binID == "" {
		return fmt.Errorf("binID не указан")
	}
	URL := c.baseUrl + "/" + binID

	req, err := http.NewRequest(http.MethodDelete, URL, nil)
	if err != nil {
		return fmt.Errorf("ошибка создания запроса: %w", err)
	}
	c.setHeders(req, false)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("соединение: %w", err)
	}
	defer resp.Body.Close()

	if err := c.checkResponse(resp); err != nil {
		return err
	}
	return nil
}
func (c *Client) apiKey() string {
	return c.cfg.Key
}

func (c *Client) setHeders(req *http.Request, contentType bool) {
	req.Header.Set("X-Master-Key", c.apiKey())
	if contentType {
		req.Header.Set("Content-Type", "application/json")
	}
}

func (c *Client) checkResponse(resp *http.Response) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("чтение тела ответа: %w", err)
	}

	switch resp.StatusCode {
	case 200, 201, 204:
		return nil
	case 401:
		return fmt.Errorf("неверный API ключ: (401)")
	case 404:
		return fmt.Errorf("хранилище не найдено: bin с таким ID не существует.")
	case 403:
		return fmt.Errorf("доступ запрещён")
	case 429:
		return fmt.Errorf("слишком много запросов, подождите")
	default:
		if resp.StatusCode >= 500 {
			return fmt.Errorf("ошибка сервера JsonBin (%d)", resp.StatusCode)
		}
		return fmt.Errorf("ошибка: %s", string(body))
	}
}

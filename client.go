package smsactivate

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	http.Client
	apiKey string
}

func New(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}

func (c *Client) Balance() (float64, error) {
	params := url.Values{
		"api_key": {c.apiKey},
		"action":  {"getBalance"},
	}
	res, err := c.Get("https://api.sms-activate.io/stubs/handler_api.php?" + params.Encode())
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	if !strings.HasPrefix(string(body), "ACCESS_BALANCE:") {
		return 0, fmt.Errorf("smsactivate: %s", body)
	}

	return strconv.ParseFloat(strings.TrimPrefix(string(body), "ACCESS_BALANCE:"), 64)
}

func (c *Client) GetNumber(service Service, country Country, maxPrice ...float64) (string, string, error) {
	params := url.Values{
		"api_key": {c.apiKey},
		"action":  {"getNumberV2"},
		"service": {string(service)},
		"country": {strconv.Itoa(int(country))},
	}
	if len(maxPrice) > 0 {
		params.Set("maxPrice", fmt.Sprintf("%.2f", maxPrice[0]))
	}
	res, err := c.Get("https://api.sms-activate.io/stubs/handler_api.php?" + params.Encode())
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", "", err
	}

	var data NumberData
	if err := json.Unmarshal(body, &data); err != nil {
		return "", "", err
	}

	if data.Error != nil {
		return "", "", fmt.Errorf("smsactivate: %s", data.Error.Msg)
	}

	return data.ActivationID, data.PhoneNumber, nil
}

func (c *Client) Wait(ctx context.Context, id string) (string, error) {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-ticker.C:
			params := url.Values{
				"api_key": {c.apiKey},
				"action":  {"getStatus"},
				"id":      {id},
			}
			res, err := c.Get("https://api.sms-activate.io/stubs/handler_api.php?" + params.Encode())
			if err != nil {
				return "", err
			}
			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			if err != nil {
				return "", err
			}

			parts := strings.Split(string(body), ":")
			if len(parts) == 1 {
				switch parts[0] {
				case "STATUS_WAIT_CODE":
					continue
				default:
					return "", fmt.Errorf("smsactivate: %s", body)
				}
			}

			return parts[1], nil
		}
	}
}

func (c *Client) Done(id string) error {
	params := url.Values{
		"api_key": {c.apiKey},
		"action":  {"setStatus"},
		"id":      {id},
		"status":  {"6"},
	}
	res, err := c.Get("https://api.sms-activate.io/stubs/handler_api.php?" + params.Encode())
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	switch string(body) {
	case "ACCESS_ACTIVATION":
		return nil
	default:
		return fmt.Errorf("smsactivate: %s", body)
	}
}

func (c *Client) Cancel(id string) error {
	params := url.Values{
		"api_key": {c.apiKey},
		"action":  {"setStatus"},
		"id":      {id},
		"status":  {"8"},
	}
	res, err := c.Get("https://api.sms-activate.io/stubs/handler_api.php?" + params.Encode())
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	switch string(body) {
	case "ACCESS_CANCEL":
		return nil
	default:
		return fmt.Errorf("smsactivate: %s", body)
	}
}

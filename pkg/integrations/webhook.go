package integrations

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func TriggerWebhook(
	url string,
	payload ...map[string]interface{},
) (string, error) {
	var data []byte
	var err error
	var response *http.Response

	if len(payload) > 0 {
		data, err = json.Marshal(payload[0])
		if err != nil {
			return "", err
		}
	}

	response, err = http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK {
		return "", errors.New(string(body))
	}

	return string(body), nil
}

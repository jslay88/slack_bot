package integrations

import (
	"fmt"
)

type JenkinsWebhookConfig struct {
	BaseURL      string
	WebhookToken string
}

func TriggerJenkinsWebhook(
	jobName string,
	config JenkinsWebhookConfig,
) (string, error) {
	url := fmt.Sprintf("%s/job/%s/build?token=%s", config.BaseURL, jobName, config.WebhookToken)
	response, err := TriggerWebhook(url)
	if err != nil {
		return "", err
	}
	return response, nil
}

func TriggerJenkinsGenericWebhook(
	config JenkinsWebhookConfig,
	payload ...map[string]interface{},
) (string, error) {
	url := fmt.Sprintf("%s/jenkins/generic-webhook-trigger/invoke?token=%s", config.BaseURL, config.WebhookToken)

	var response string
	var err error

	if len(payload) > 0 {
		response, err = TriggerWebhook(url, payload[0])
	} else {
		response, err = TriggerWebhook(url)
	}
	if err != nil {
		return "", err
	}
	return response, nil
}

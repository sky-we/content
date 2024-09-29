package services

import (
	"content-system/internal/config"
	"fmt"
	"net/http"
	"strings"
)

func (app *CmsApp) startContentFlow(id int64, config *config.FlowServiceClientConfig) error {
	host := config.Host
	port := config.Port
	flowName := config.FlowName
	body := fmt.Sprintf(`{"input": %d}`, id)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%v:%v/%v", host, port, flowName), strings.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	_, reqErr := client.Do(req)
	if reqErr != nil {
		return err
	}
	return nil
}

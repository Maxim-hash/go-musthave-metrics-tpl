package agent

import (
	"fmt"
	"net/http"
)

type Sender struct {
	ServerAddr string
	Client     *http.Client
}

func NewSender(serverAddr string) *Sender {
	return &Sender{
		ServerAddr: serverAddr,
		Client:     &http.Client{},
	}
}

func (s *Sender) SendMetric(metricType, metricName, metricValue string) error {
	r := fmt.Sprintf("http://%s/update/%s/%s/%s", s.ServerAddr, metricType, metricName, metricValue)
	req, err := http.NewRequest(http.MethodPost, r, nil)
	if err != nil {
		return fmt.Errorf("create request failed: %w", err)
	}
	req.Header.Set("Content-Type", "text/plain")
	resp, err := s.Client.Do(req)
	if err != nil {
		return fmt.Errorf("send metric failed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned %d: %w", resp.StatusCode, err)
	}

	return nil
}

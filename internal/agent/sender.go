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
	r := fmt.Sprintf("%s/update/%s/%s/%s", s.ServerAddr, metricType, metricName, metricValue)
	req, err := http.NewRequest(http.MethodPost, r, nil)
	if err != nil {
		return err
	}
	_, err = s.Client.Do(req)
	return err
}

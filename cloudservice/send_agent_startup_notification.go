package cloudservice

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/docktermj/go-logger/logger"
)

func (self CloudService) SendStartupNotification(nodeID string) error {

	message := map[string]string{
		"nodeid": nodeID,
	}
	body, err := json.Marshal(message)
	if err != nil {
		logger.Errorf("Failed encoding message: %v", err)
		return err
	}

	_, err = http.Post(self.cloudServiceURL+"/v1/agent-startup", "application/json", bytes.NewBuffer(body))
	if err != nil {
		logger.Errorf("Failed contacting cloud service: %v", err)
		return err
	}

	logger.Infof("Contacted cloud service at %s", self.cloudServiceURL)
	return nil
}

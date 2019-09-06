package cloudservice

import (
	"encoding/json"

	"github.com/docktermj/go-logger/logger"
	"github.com/drauschenbach/megalithicd/propertykey"
	"github.com/gorilla/websocket"
)

func (self CloudService) SendStartupNotification() error {

	nodeID, err := self.propertiesDAO.Get(propertykey.NodeID)
	if err != nil {
		logger.Fatalf("Failed looking up node id: %v", err)
		return nil
	}

	message := map[string]string{
		"nodeid": nodeID,
	}
	body, err := json.Marshal(message)
	if err != nil {
		logger.Errorf("Failed encoding message: %v", err)
		return err
	}

	if err := self.conn.WriteMessage(websocket.TextMessage, body); err != nil {
		logger.Errorf("Failed contacting cloud service: %v", err)
		return err
	}

	logger.Infof("Contacted cloud service")
	return nil
}

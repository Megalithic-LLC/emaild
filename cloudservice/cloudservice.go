package cloudservice

import (
	"os"

	"github.com/docktermj/go-logger/logger"
	"github.com/drauschenbach/megalithicd/dao"
	"github.com/drauschenbach/megalithicd/propertykey"
)

type CloudService struct {
	cloudServiceURL string
}

func New(propertiesDAO dao.PropertiesDAO) *CloudService {
	cloudServiceURL := os.Getenv("CLOUDSERVICE_URL")
	if cloudServiceURL == "" {
		cloudServiceURL = "http://localhost:3000"
	}

	nodeID, err := propertiesDAO.Get(propertykey.NodeID)
	if err != nil {
		logger.Fatalf("Failed looking up node id: %v", err)
		return nil
	}

	self := CloudService{
		cloudServiceURL: cloudServiceURL,
	}

	if err := self.SendStartupNotification(nodeID); err != nil {
		logger.Fatalf("Failed contacting cloud service: %v", err)
		return nil
	}

	return &self
}

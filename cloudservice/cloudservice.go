package cloudservice

import (
	"net/url"
	"os"
	"time"

	"github.com/docktermj/go-logger/logger"
	"github.com/drauschenbach/megalithicd/dao"
	"github.com/gorilla/websocket"
)

const CLOUDSERVICE_URL = "CLOUDSERVICE_URL"

type CloudService struct {
	cloudServiceURL url.URL
	conn            *websocket.Conn
	propertiesDAO   dao.PropertiesDAO
}

func New(propertiesDAO dao.PropertiesDAO) *CloudService {
	cloudServiceURL := os.Getenv(CLOUDSERVICE_URL)
	if cloudServiceURL == "" {
		cloudServiceURL = "http://localhost:3000"
	}
	parsedURL, err := url.Parse(cloudServiceURL)
	if err != nil {
		logger.Fatalf("Malformed %s", CLOUDSERVICE_URL)
		return nil
	}

	self := CloudService{
		cloudServiceURL: url.URL{Scheme: "ws", Host: parsedURL.Host, Path: "/v1/agentstream"},
		propertiesDAO:   propertiesDAO,
	}

	go self.dialer()

	return &self
}

func (self *CloudService) dialer() {
	for {
		if self.conn != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		logger.Debugf("Connecting to %s", self.cloudServiceURL.String())
		conn, _, err := websocket.DefaultDialer.Dial(self.cloudServiceURL.String(), nil)
		if err != nil {
			self.conn = nil
			time.Sleep(15 * time.Second)
			continue
		}
		logger.Debugf("Connected to %s", self.cloudServiceURL.String())
		self.conn = conn

		if err := self.SendStartupNotification(); err != nil {
			logger.Warnf("Failed contacting cloud service: %v", err)
		}
	}
}

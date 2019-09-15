package syncservice

import (
	"github.com/docktermj/go-logger/logger"
	"github.com/drauschenbach/megalithicd/dao"
)

type SyncService struct {
	propertiesDAO dao.PropertiesDAO
}

func New(propertiesDAO dao.PropertiesDAO) *SyncService {
	self := SyncService{
		propertiesDAO: propertiesDAO,
	}
	if err := self.Start(); err != nil {
		logger.Fatalf("Failed starting sync service: %v", err)
		return nil
	}
	logger.Infof("Sync service started")
	return &self
}

func (self *SyncService) Close() {
	logger.Tracef("SyncService:Close()")
}

//func (self *SyncService) reader() {
//}

func (self *SyncService) Start() error {
	logger.Tracef("SyncService:Start()")
	return nil
}

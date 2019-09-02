package main

import (
	"github.com/drauschenbach/megalithicd/propertykey"
	"github.com/rs/xid"
)

func getOrGenerateNodeID() (string, error) {
	return propertiesDAO.SetIfKeyNotExists(propertykey.NodeID, xid.New().String())
}

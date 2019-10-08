package main

import (
	"github.com/on-prem-net/emaild/propertykey"
	"github.com/rs/xid"
)

func getOrGenerateNodeID() (string, error) {
	return propertiesDAO.SetIfKeyNotExists(propertykey.NodeID, xid.New().String())
}

package main

import (
	"github.com/Megalithic-LLC/on-prem-emaild/propertykey"
	"github.com/rs/xid"
)

func getOrGenerateNodeID() (string, error) {
	return propertiesDAO.SetIfKeyNotExists(propertykey.NodeID, xid.New().String())
}

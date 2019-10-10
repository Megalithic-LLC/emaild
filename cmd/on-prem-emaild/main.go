package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/docktermj/go-logger/logger"
)

var ()

func init() {
	logger.SetLevel(logger.LevelTrace)
}

func main() {
	DefineDependencies()
	graph.ResolveAll()

	// If a unique id does not yet exist, create it now
	nodeid, err := getOrGenerateNodeID()
	if err != nil {
		logger.Fatalf("Failed assigning node id: %v", err)
	}

	logger.Info("On-Prem Email Server started")
	logger.Infof("Node id is %s", nodeid)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Wait for shutdown
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	_ = ctx
	imapEndpoint.Shutdown()
	smtpEndpoint.Shutdown()
	submissionEndpoint.Shutdown()
	logger.Infof("Shutting down")
	if db != nil {
		db.Close()
		logger.Infof("Closed database")
	}
	os.Exit(0)

}

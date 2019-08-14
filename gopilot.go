package main

import (
	"os/signal"
	"os"
	"github.com/sirupsen/logrus"
	"flag"
	"gitlab.com/gopilot/lib/config"
	"gitlab.com/gopilot/lib/mynodename"
	"gitlab.com/gopilot/gopilotd/gopilotd"
	log_prefixed "github.com/chappjc/logrus-prefix"
)

type cmdLineParamsT struct {
	debug             bool
	SocketBusFileName string
}

var cmdLineParams cmdLineParamsT

func main() {

	// ########################## Command line parse ##########################
	flag.BoolVar(&cmdLineParams.debug, "v", false, "Enable debug")
	flag.StringVar(&cmdLineParams.SocketBusFileName, "socket", "/tmp/gopilotd.sock", "Socket of gopilotd")
	mynodename.ParseCmdLine()
	config.ParseCmdLine()
	flag.Parse()

	// ########################## logging ##########################
	logrus.SetFormatter(new(log_prefixed.TextFormatter))
	if cmdLineParams.debug == true {
		logrus.SetLevel(logrus.DebugLevel)
	}

	


	// ########################## Init ##########################
	mynodename.Init()
	config.Init()
	config.Read()

	// ########################## The Bus ##########################
	gopilotd.Init()

	// atexit
	go func( ) {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		
		// Block until a signal is received.
		s := <-c
		logrus.Debugf("Got signal '%s'", s)
		if err := os.RemoveAll(cmdLineParams.SocketBusFileName); err != nil {
			logrus.Error(err)
		}
		os.Exit(0)
    }()


	gopilotd.Serve(cmdLineParams.SocketBusFileName)


}

package main

import (
	"24/challenges/day1"
	"flag"
	"strconv"
	"time"

	"github.com/charmbracelet/log"
)

func main() {
	log.SetReportCaller(true)
	log.SetReportTimestamp(false)
	log.SetFormatter(log.TextFormatter)

	var mode string
	verbosePtr := flag.Bool("v", false, "Show debug logs")
	flag.Parse()
	if *verbosePtr {
		log.SetLevel(log.DebugLevel)
	}

	if len(flag.Args()) < 1 {
		log.Fatalf("Need to specify challenge date, got args %v", flag.Args())
	}
	chal, err := strconv.Atoi(flag.Args()[0])
	if err != nil {
		log.Fatal("Challenge date must be an int", "invalid challenge", chal, "error", err)
	}

	if len(flag.Args()) != 2 {
		mode = "test"
	} else {
		mode = "final"
	}
	startTime := time.Now()
	switch chal {
	case 1:
		day1.Sol(mode)
	}
	executionTime := time.Since(startTime)
	defer log.Infof("Finished executing in %s", executionTime)
}
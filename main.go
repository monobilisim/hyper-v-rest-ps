package main

import (
	"flag"
	"os"
	"time"

	"hyper-v-rest-ps/hyperv"
	"hyper-v-rest-ps/rest"
	"hyper-v-rest-ps/utilities"

	"github.com/kardianos/service"
)

var logger service.Logger
var log = utilities.Log

type program struct {
	exit chan struct{}
}

func (p *program) Start(s service.Service) error {
	if service.Interactive() {
		logger.Info("Running in interactive terminal.")
	} else {
		logger.Info("Running as service.")
	}
	p.exit = make(chan struct{})

	go p.run()

	return nil
}

func (p *program) run() {
	c := utilities.ParseConfig()
	rest.StartServer(c.Port, "1.3.1")
}

func (p *program) Stop(s service.Service) error {
	if service.Interactive() {
		os.Exit(0)
	} else {
		logger.Info("Stopping service...")
		close(p.exit)
	}
	return nil
}

func main() {
	svcFlag := flag.String("service", "", "Control the system service.")
	flag.Parse()

	utilities.Init()
	utilities.SetupLogger()
	utilities.Wg.Add(1)
	go func() {
		time.Sleep(10 * time.Second)
		log.Info("Application started.")
		utilities.Wg.Done()
	}()

	hyperv.Init()

	go func() {
		for {
			time.Sleep(660 * time.Second)
			hyperv.Refresh()
			log.Info("Hyper-V module reinitialized.")
		}
	}()

	go func() {
		for {
			time.Sleep(1919 * time.Second)
			utilities.RefreshShellQueue()
			log.Info("Shell queue reinitialized.")
		}
	}()

	options := make(service.KeyValue)
	options["Restart"] = "on-success"
	options["SuccessExitStatus"] = "1 2 8 SIGKILL"
	svcConfig := &service.Config{
		Name:        "hyper-v-rest-ps",
		DisplayName: "Hyper-V REST PowerShell",
		Description: "Simple REST service for some Hyper-V features.",
		Option:      options,
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	errs := make(chan error, 5)
	logger, err = s.Logger(errs)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			err := <-errs
			if err != nil {
				log.Print(err)
			}
		}
	}()

	if len(*svcFlag) != 0 {
		err := service.Control(s, *svcFlag)
		if err != nil {
			log.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}
		return
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}

	select {}
}

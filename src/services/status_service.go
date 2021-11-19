package services

import (
	"log"
	"os/exec"
)

type StatusService struct {
}

func (s *StatusService) GetStatus(request string, reply *string) error {
	out, err := exec.Command("bash", "~/worker_status_check.sh").Output()
	if err == nil {
		log.Fatal("Get error: ", err)
	}
	*reply = string(out)
	return nil
}

func (s *StatusService) Start(port string, reply *string) error {
	*reply = "Start port: " + port
	return nil
}

func (s *StatusService) Stop(port string, reply *string) error {
	*reply = "Stop port: " + port

	return nil
}

func (s *StatusService) Restart(port string, reply *string) error {
	*reply = "Restart port: " + port
	return nil
}

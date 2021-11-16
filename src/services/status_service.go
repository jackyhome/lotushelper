package services

type StatusService struct {
}

func (s *StatusService) GetStatus(request string, reply *string) error {
	*reply = "Status: " + request
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

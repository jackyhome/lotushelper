package interfaces

import "net/rpc"

const STATUS_SERVICE_NAME = "github.com/jackyhome/StatusService"

type StatusServiceClient struct {
	*rpc.Client
}

func RegisterStatusService(svc StatusServiceInterface) error {
	return rpc.RegisterName(STATUS_SERVICE_NAME, svc)
}

type StatusServiceInterface interface {
	GetStatus(request string, reply *string) error
	Stop(port string, reply *string) error
	Start(port string, reply *string) error
	Restart(port string, reply *string) error
}

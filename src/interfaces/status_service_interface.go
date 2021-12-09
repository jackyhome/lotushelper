package interfaces

import (
	"net/rpc"

	"github.com/jackyhome/lotushelper/src/models"
)

const STATUS_SERVICE_NAME = "github.com/jackyhome/StatusService"

type StatusServiceClient struct {
	*rpc.Client
}

func RegisterStatusService(svc StatusServiceInterface) error {
	return rpc.RegisterName(STATUS_SERVICE_NAME, svc)
}

type LtStatus struct {
	ServerName   string
	ServerType   string
	ServerHost   string
	DiskInfo     string
	GpuStatus    models.StatusCode
	InstanceList []models.LtInstance
}
type StatusServiceInterface interface {
	GetStatus(request string, reply *string) error
	UpdateStatus(request LtStatus, reply *string) error
	Stop(port string, reply *string) error
	Start(port string, reply *string) error
	Restart(port string, reply *string) error
}

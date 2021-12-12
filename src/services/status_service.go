package services

import (
	"log"
	"time"

	//"os/exec"

	"github.com/jackyhome/lotushelper/src/interfaces"
	"github.com/jackyhome/lotushelper/src/models"
)

type StatusService struct {
}

func (s *StatusService) GetStatus(request string, reply *string) error {
	//out, err := exec.Command("bash", "/Users/jacky/worker_status_check.sh").CombinedOutput()
	// if err != nil {
	// 	log.Fatal("Get error: ", err)
	// }
	// log.Println(string(out))
	status := models.GetLtStatus()
	serverStatus, ok := status.GetServerStatus(request)
	if ok {
		*reply = serverStatus.String()
	} else {
		*reply = status.String()
	}
	return nil
}

func (s *StatusService) UpdateStatus(request interfaces.LtStatus, reply *string) error {
	log.Println("Calling update with: " + request.ServerName)
	sStatus := *(models.GetLtStatus())
	serverStatus, ok := sStatus.GetServerStatus(request.ServerName)

	if !ok {
		serverStatus = &models.ServerStatus{
			Server: &models.LtServer{
				ServerType: models.ServerType(request.ServerType),
				ServerName: request.ServerName,
				ServerHost: request.ServerHost,
			},
			GpuStatus:      request.GpuStatus,
			DiskInfo:       request.DiskInfo,
			Instances:      map[string]models.LtInstance{},
			InstanceStatus: map[string]*models.InstanceStatus{},
			UpdateTime:     time.Now(),
		}
		for _, instance := range request.InstanceList {
			serverStatus.InstanceStatus[instance.InstanceName] = &models.InstanceStatus{
				StatusCode: models.ON,
				TaskCount:  0,
			}
			serverStatus.Instances[instance.InstanceName] = instance
		}
		sStatus.StatusMap[request.ServerName] = serverStatus
	} else {
		serverStatus.StatusLck.Lock()
		defer serverStatus.StatusLck.Unlock()

		for _, v := range serverStatus.InstanceStatus {
			v.StatusCode = models.OFF
		}
		serverStatus.Server.ServerType = models.ServerType(request.ServerType)
		serverStatus.GpuStatus = request.GpuStatus
		serverStatus.DiskInfo = request.DiskInfo
		serverStatus.UpdateTime = time.Now()
		for _, instance := range request.InstanceList {
			serverStatus.InstanceStatus[instance.InstanceName].StatusCode = models.ON
		}

	}
	*reply = serverStatus.String()
	return nil
}

func (s *StatusService) Start(request string, reply *string) error {
	*reply = "Start port: " + request
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

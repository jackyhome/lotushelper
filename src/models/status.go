package models

import (
	"encoding/json"
	"sync"
)

type ServerType string

const (
	MMiner  ServerType = "Main Miner"
	WMiner  ServerType = "Winning Miner"
	MNode   ServerType = "Main Node"
	BNode   ServerType = "Backup Node"
	CWorker ServerType = "Commit Worker"
	PWorker ServerType = "PreCommit Worker"
)

type StatusCode int

const (
	ON  StatusCode = 1
	OFF StatusCode = 0
)

type LtServer struct {
	ServerType  ServerType
	ServerName  string
	ServerHost  string
	LtInstances []LtInstance
}
type LtInstance struct {
	InstanceName string
	StartCmd     string
}
type InstanceStatus struct {
	StatusCode
	TaskCount int
}
type ServerStatus struct {
	Server         *LtServer
	StatusLck      sync.Mutex
	GpuStatus      StatusCode
	DiskInfo       string
	Instances      map[string]LtInstance
	InstanceStatus map[string]*InstanceStatus
}

var instantiated *ServerStatus
var once sync.Once

func GetStatus() *ServerStatus {
	once.Do(func() {
		instantiated = &ServerStatus{
			Server: &LtServer{
				ServerType: MMiner,
				ServerName: "testName",
			},
		}
	})
	return instantiated
}

func (s *ServerStatus) String() string {
	val, _ := json.Marshal(s)
	return string(val)
}

var ltStatus *LtStatus

type LtStatus struct {
	StatusMap map[string]*ServerStatus
}

func (s *LtStatus) String() string {
	val, _ := json.Marshal(s)
	return string(val)
}
func (s *LtStatus) GetServerStatus(name string) (*ServerStatus, bool) {
	serverStatus, ok := s.StatusMap[name]
	return serverStatus, ok
}
func GetLtStatus() *LtStatus {
	once.Do(func() {
		ltStatus = &LtStatus{
			StatusMap: map[string]*ServerStatus{},
		}
	})
	return ltStatus
}

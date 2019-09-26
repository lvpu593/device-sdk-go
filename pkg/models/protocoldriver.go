package models

import (
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	contract "github.com/edgexfoundry/go-mod-core-contracts/models"
)

type ProtocolDriver interface {
	Initialize(lc logger.LoggingClient, asyncCh chan<- *AsyncValues) error

	HandlerReadCommands(deviceName string, protocols map[string]contract.ProtocolProperties, reqs []CommandRequest) ([]*CommandValue, error)

	HandlerWriteCommands(deviceName string, protocols map[string]contract.ProtocolProperties, reqs []CommandRequest, params []*CommandValue) error

	Stop(force bool) error

	AddDevice(deviceName string, procotols map[string]contract.ProtocolProperties, reqs []CommandRequest, adminState contract.AdminState) error

	UpdateDevice(deviceName string, procotols map[string]contract.ProtocolProperties, reqs []CommandRequest, adminState contract.AdminState) error

	RemoveDevice(deviceName string, procotols map[string]contract.ProtocolProperties, reqs []CommandRequest) error
}

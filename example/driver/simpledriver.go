package driver

import (
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	"github.com/lvpu593/device-sdk-go/pkg/models"
)

type SimpleDriver struct {
	lc      logger.LoggingClient
	asyncCh chan<- models.AsyncValues
}

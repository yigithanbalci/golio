package val

import (
	"github.com/KnutZuidema/golio/internal"
	log "github.com/sirupsen/logrus"
)

type StatusClient struct {
	c *internal.Client
}

func (cc *StatusClient) GetPlatformData() (*PlatformData, error) {
	logger := cc.logger().WithField("method", "GetPlatformData")
	var platformData *PlatformData
	if err := cc.c.GetInto(endpointGetPlatformData, &platformData); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return platformData, nil
}

func (cc *StatusClient) logger() log.FieldLogger {
	return cc.c.Logger().WithField("category", "status")
}

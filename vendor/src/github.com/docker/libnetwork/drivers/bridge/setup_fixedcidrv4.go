package bridge

import (
	log "github.com/Sirupsen/logrus"
	"github.com/docker/libnetwork/ipallocator"
)

func setupFixedCIDRv4(config *networkConfiguration, i *bridgeInterface) error {
	addrv4, _, err := i.addresses()
	if err != nil {
		return err
	}
	if ipAllocator[config.VlanId] == nil {
		ipAllocator[config.VlanId] = ipallocator.New()
	}

	log.Debugf("Using IPv4 subnet: %v", config.FixedCIDR)
	if err := ipAllocator[config.VlanId].RegisterSubnet(addrv4.IPNet, config.FixedCIDR[config.VlanId]); err != nil {
		return &FixedCIDRv4Error{Subnet: config.FixedCIDR[config.VlanId], Net: addrv4.IPNet, Err: err}
	}

	return nil
}

package bind

import (
	"fmt"
	"net"

	"github.com/inverse-inc/packetfence/go/coredns/core/dnsserver"
	"github.com/inverse-inc/packetfence/go/coredns/plugin"

	"github.com/inverse-inc/packetfence/go/caddy/caddy"
)

func setupBind(c *caddy.Controller) error {
	config := dnsserver.GetConfig(c)
	for c.Next() {
		if !c.Args(&config.ListenHost) {
			return plugin.Error("bind", c.ArgErr())
		}
	}
	if net.ParseIP(config.ListenHost) == nil {
		return plugin.Error("bind", fmt.Errorf("not a valid IP address: %s", config.ListenHost))
	}
	return nil
}

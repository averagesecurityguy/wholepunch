package wp

import (
	"fmt"

	"github.com/tomsteele/wholepunch/pkg/tls"
)

// BeaconTLSConnect uses TCP to connect to a TLS listener. No data is sent over the channel.
type BeaconTLSConnect struct {
	ServerAddr string
}

// Name returns the name of the module.
func (b *BeaconTLSConnect) Name() string {
	return "tls-tcp-connect"
}

// Destination returns the server that was connected to.
func (b *BeaconTLSConnect) Destination() string {
	return b.ServerAddr
}

// Success returns a formatted string indicating a successfull connection.
func (b *BeaconTLSConnect) Success() string {
	return fmt.Sprintf("The agent was allowed to reach %s using TLS over TCP.", b.ServerAddr)
}

// Setup is used to initilize instance variables from BeaconOptions.
func (b *BeaconTLSConnect) Setup(o *BeaconOptions) error {
	b.ServerAddr = o.DestinationServerAddress
	return nil
}

// Send initiates the TCP/TLS connection.
func (b *BeaconTLSConnect) Send() (bool, error) {
	return tls.DialInsecureTCP(b.ServerAddr)
}

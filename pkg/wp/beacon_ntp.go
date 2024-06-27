package wp

import (
	"fmt"

	"github.com/tomsteele/wholepunch/pkg/ntp"
)

// BeaconNTP creates a NTP connection to a server.
type BeaconNTP struct {
	ServerAddr string
	Timeout    int
}

// Name returns the name of the beacon operation.
func (b *BeaconNTP) Name() string {
	return "ntp"
}

// Destination returns the server that was connected to.
func (b *BeaconNTP) Destination() string {
	return fmt.Sprintf("ntp://%s", b.ServerAddr)
}

// Success returns a formatted string indicating a successful connection.
func (b *BeaconNTP) Success() string {
	return fmt.Sprintf("The agent was allowed to reach %s using NTP.", b.ServerAddr)
}

// Setup is used to initilize instance variables from BeaconOptions.
func (b *BeaconNTP) Setup(o *BeaconOptions) error {
	b.ServerAddr = o.DestinationServerAddress
	return nil
}

// Send creates a NTP connection.
func (b *BeaconNTP) Send() (bool, error) {
	return ntp.GetCurrentTime(b.ServerAddr)
}

package wp

import (
	"fmt"
	"net"
	"time"
)

// BeaconTCP creates a TCP connection to a server.
type BeaconTCP struct {
	ServerAddr string
	Timeout    int
}

// Name returns the name of the beacon operation.
func (b *BeaconTCP) Name() string {
	return "tcp"
}

// Destination returns the server that was connected to.
func (b *BeaconTCP) Destination() string {
	return fmt.Sprintf("tcp://%s", b.ServerAddr)
}

// Success returns a formatted string indicating a successfull connection.
func (b *BeaconTCP) Success() string {
	return fmt.Sprintf("The agent was allowed to reach %s using TCP.", b.ServerAddr)
}

// Setup is used to initilize instance variables from BeaconOptions.
func (b *BeaconTCP) Setup(o *BeaconOptions) error {
	b.ServerAddr = o.DestinationServerAddress
	return nil
}

// Send creates a TCP connection.
func (b *BeaconTCP) Send() (bool, error) {
	conn, err := net.DialTimeout("tcp", b.ServerAddr, time.Duration(b.Timeout)*time.Millisecond)
	if err != nil {
		return false, err
	}
	conn.Close()
	return true, nil
}

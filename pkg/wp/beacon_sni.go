package wp

import (
	"fmt"

	"github.com/tomsteele/wholepunch/pkg/tls"
)

// BeaconSpoofSNI creates a TLS connection with a ServerName in the Client Hello portion
// of the TLS handshake set to your provided hostname.
type BeaconSpoofSNI struct {
	ServerName string
	ServerAddr string
}

// Name returns the name of the beacon operation.
func (b *BeaconSpoofSNI) Name() string {
	return "tls-sni-spoof"
}

// Destination returns the server that was connected to
func (b *BeaconSpoofSNI) Destination() string {
	return b.ServerAddr
}

// Success returns a formatted string indicating a successfull connection.
func (b *BeaconSpoofSNI) Success() string {
	return fmt.Sprintf("The agent was allowed to reach %s using TLS/TCP. The SNI name of %s was used to bypass egress controls.", b.ServerAddr, b.ServerName)
}

// Setup is used to initilize instance variables from BeaconOptions.
func (b *BeaconSpoofSNI) Setup(o *BeaconOptions) error {
	b.ServerAddr = o.DestinationServerAddress
	return nil
}

// Send sends a TLS request.
func (b *BeaconSpoofSNI) Send() (bool, error) {
	return tls.SendSpoofSNI(b.ServerName, b.ServerAddr)
}

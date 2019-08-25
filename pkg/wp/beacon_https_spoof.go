package wp

import (
	"fmt"

	"github.com/tomsteele/wholepunch/pkg/http"
)

// BeaconHTTPSGetSpoofHostHeader ...
type BeaconHTTPSGetSpoofHostHeader struct {
	HostHeaderName string
	UserAgent      string
	ServerAddr     string
	Path           string
	ServerURL      string
}

// Name returns the name of the module.
func (b *BeaconHTTPSGetSpoofHostHeader) Name() string {
	return "https-host-spoof"
}

// Destination returns the server that was connected to
func (b *BeaconHTTPSGetSpoofHostHeader) Destination() string {
	return b.ServerURL
}

// Success returns a formatted string indicating a successfull connection.
func (b *BeaconHTTPSGetSpoofHostHeader) Success() string {
	return fmt.Sprintf("The agent was allowed to communicate with %s over HTTP(TLS) using the Host header %s", b.ServerAddr, b.HostHeaderName)
}

// Setup is used to initilize instance variables from BeaconOptions.
func (b *BeaconHTTPSGetSpoofHostHeader) Setup(o *BeaconOptions) error {
	b.ServerAddr = o.DestinationServerAddress
	b.ServerURL = fmt.Sprintf("https://%s%s", b.ServerAddr, b.Path)
	return nil
}

// Send initiates the TCP/TLS connection.
func (b *BeaconHTTPSGetSpoofHostHeader) Send() (bool, error) {
	return http.TLSGetSpoofHostHeader(b.ServerURL, b.HostHeaderName, b.UserAgent)
}

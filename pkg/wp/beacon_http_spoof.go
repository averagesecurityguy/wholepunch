package wp

import (
	"fmt"

	"github.com/tomsteele/wholepunch/pkg/http"
)

// BeaconHTTPGetSpoofHostHeader sends an HTTP GET request to
// ServerAddr using a modified Host header.
type BeaconHTTPGetSpoofHostHeader struct {
	HostHeaderName string
	UserAgent      string
	ServerAddr     string
	Path           string
	ServerURL      string
}

// Name returns the name of the module.
func (b *BeaconHTTPGetSpoofHostHeader) Name() string {
	return "http-get-spoof-host-header"
}

// Destination returns the server that was connected to
func (b *BeaconHTTPGetSpoofHostHeader) Destination() string {
	return b.ServerURL
}

// Success returns a formatted string indicating a successfull connection.
func (b *BeaconHTTPGetSpoofHostHeader) Success() string {
	return fmt.Sprintf("The agent was allowed to communicate with %s over HTTP using the Host header %s", b.ServerAddr, b.HostHeaderName)
}

// Setup is used to initilize instance variables from BeaconOptions.
func (b *BeaconHTTPGetSpoofHostHeader) Setup(o *BeaconOptions) error {
	b.ServerAddr = o.DestinationServerAddress
	b.ServerURL = fmt.Sprintf("http://%s%s", b.ServerAddr, b.Path)
	return nil
}

// Send initiates the TCP/TLS connection.
func (b *BeaconHTTPGetSpoofHostHeader) Send() (bool, error) {
	return http.GetSpoofHostHeader(b.ServerURL, b.HostHeaderName, b.UserAgent)
}

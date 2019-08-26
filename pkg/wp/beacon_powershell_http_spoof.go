package wp

import (
	"fmt"

	"github.com/tomsteele/wholepunch/pkg/powershell"
)

// BeaconPowershellHTTPGetSpoofHostHeader sends an HTTP GET request to
// ServerAddr using a modified Host header.
type BeaconPowershellHTTPGetSpoofHostHeader struct {
	HostHeaderName string
	UserAgent      string
	ServerAddr     string
	Path           string
	ServerURL      string
}

// Name returns the name of the module.
func (b *BeaconPowershellHTTPGetSpoofHostHeader) Name() string {
	return "powershell-http-get-spoof-host-header"
}

// Destination returns the server that was connected to
func (b *BeaconPowershellHTTPGetSpoofHostHeader) Destination() string {
	return b.ServerURL
}

// Success returns a formatted string indicating a successfull connection.
func (b *BeaconPowershellHTTPGetSpoofHostHeader) Success() string {
	return fmt.Sprintf("The agent was allowed to communicate with %s over HTTP with Powershell using the Host header %s", b.ServerAddr, b.HostHeaderName)
}

// Setup is used to initilize instance variables from BeaconPowershellOptions.
func (b *BeaconPowershellHTTPGetSpoofHostHeader) Setup(o *BeaconOptions) error {
	b.ServerAddr = o.DestinationServerAddress
	b.ServerURL = fmt.Sprintf("http://%s%s", b.ServerAddr, b.Path)
	return nil
}

// Send initiates the TCP/TLS connection.
func (b *BeaconPowershellHTTPGetSpoofHostHeader) Send() (bool, error) {
	script := fmt.Sprintf("Invoke-WebRequest -Headers @{\"Host\"=\"%s\"; \"User-Agent\"=\"%s\"} %s", b.HostHeaderName, b.UserAgent, b.ServerURL)
	err := powershell.RunCommand(script)
	return err == nil, err
}

package wp

import (
	"fmt"

	"github.com/tomsteele/wholepunch/pkg/powershell"
)

// BeaconPowershellHTTPSGetSpoofHostHeader sends an HTTPS GET request to
// ServerAddr using a modified Host header.
type BeaconPowershellHTTPSGetSpoofHostHeader struct {
	HostHeaderName string
	UserAgent      string
	ServerAddr     string
	Path           string
	ServerURL      string
}

// Name returns the name of the module.
func (b *BeaconPowershellHTTPSGetSpoofHostHeader) Name() string {
	return "powershell-https-get-spoof-host-header"
}

// Destination returns the server that was connected to
func (b *BeaconPowershellHTTPSGetSpoofHostHeader) Destination() string {
	return b.ServerURL
}

// Success returns a formatted string indicating a successfull connection.
func (b *BeaconPowershellHTTPSGetSpoofHostHeader) Success() string {
	return fmt.Sprintf("The agent was allowed to communicate with %s over HTTPS with Powershell using the Host header %s", b.ServerAddr, b.HostHeaderName)
}

// Setup is used to initilize instance variables from BeaconPowershellOptions.
func (b *BeaconPowershellHTTPSGetSpoofHostHeader) Setup(o *BeaconOptions) error {
	b.ServerAddr = o.DestinationServerAddress
	b.ServerURL = fmt.Sprintf("https://%s%s", b.ServerAddr, b.Path)
	return nil
}

// Send initiates the TCP/TLS connection.
func (b *BeaconPowershellHTTPSGetSpoofHostHeader) Send() (bool, error) {
	script := fmt.Sprintf("Invoke-WebRequest -Headers @{\"Host\"=\"%s\"; \"User-Agent\"=\"%s\"} %s", b.HostHeaderName, b.UserAgent, b.ServerURL)
	err := powershell.RunCommand(script)
	return err == nil, err
}

package wp

import (
	"fmt"

	"github.com/tomsteele/wholepunch/pkg/http"
)

// BeaconHTTPSGet sends an HTTP GET request using TLS to a desitination server.
type BeaconHTTPSGet struct {
	UserAgent  string
	ServerAddr string
	Path       string
	ServerURL  string
}

// Name returns the name of the module.
func (b *BeaconHTTPSGet) Name() string {
	return "https-get"
}

// Destination returns the server that was connected to
func (b *BeaconHTTPSGet) Destination() string {
	return b.ServerURL
}

// Success returns a formatted string indicating a successfull connection.
func (b *BeaconHTTPSGet) Success() string {
	return fmt.Sprintf("The agent was allowed to communicate with %s over HTTP(TLS)", b.ServerAddr)
}

// Setup is used to initilize instance variables from BeaconOptions.
func (b *BeaconHTTPSGet) Setup(o *BeaconOptions) error {
	b.ServerAddr = o.DestinationServerAddress
	b.ServerURL = fmt.Sprintf("https://%s%s", b.ServerAddr, b.Path)
	return nil
}

// Send sends the HTTP GET request.
func (b *BeaconHTTPSGet) Send() (bool, error) {
	return http.TLSGet(b.ServerURL, b.UserAgent)
}

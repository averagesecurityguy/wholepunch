package wp

import (
	"fmt"

	"github.com/tomsteele/wholepunch/pkg/http"
)

// BeaconHTTPGet sends an HTTP GET request to a destination server.
type BeaconHTTPGet struct {
	UserAgent  string
	ServerAddr string
	Path       string
	ServerURL  string
}

// Name returns the name of the module.
func (b *BeaconHTTPGet) Name() string {
	return "http-get"
}

// Destination returns the server that was connected to
func (b *BeaconHTTPGet) Destination() string {
	return b.ServerURL
}

// Success returns a formatted string indicating a successfull connection.
func (b *BeaconHTTPGet) Success() string {
	return fmt.Sprintf("The agent was allowed to communicate with %s over HTTP", b.ServerAddr)
}

// Setup is used to initilize instance variables from BeaconOptions.
func (b *BeaconHTTPGet) Setup(o *BeaconOptions) error {
	b.ServerAddr = o.DestinationServerAddress
	b.ServerURL = fmt.Sprintf("http://%s%s", b.ServerAddr, b.Path)
	return nil
}

// Send sends the HTTP Get request.
func (b *BeaconHTTPGet) Send() (bool, error) {
	return http.Get(b.ServerURL, b.UserAgent)
}

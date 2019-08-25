package wp

import "fmt"

// BeaconOptions are passed during a Beacon's setup.
// These are generic options that can be used across many different tasks.
type BeaconOptions struct {
	DestinationServerAddress string
	DownloadHash             string
}

// BeaconResult holds result information for a beacon execution.
type BeaconResult struct {
	Name        string
	Destination string
	Err         error
	WasOk       bool
	Info        string
}

// Beacon is an interface for a outbound connection from an agent to a listening server.
type Beacon interface {
	Name() string
	Destination() string
	Setup(*BeaconOptions) error
	Send() (bool, error)
	Success() string
}

// RunBeacon executes a Beacon.
func RunBeacon(beacon Beacon, opts *BeaconOptions) (bool, error) {
	beacon.Setup(opts)
	return beacon.Send()
}

// MakeBeaconResult creates a BeaconResult.
func MakeBeaconResult(ok bool, err error, b Beacon) BeaconResult {
	infostr := b.Success()
	if err != nil {
		infostr = fmt.Sprintf("Error: %s.", err.Error())
	}
	return BeaconResult{
		WasOk:       ok,
		Destination: b.Destination(),
		Name:        b.Name(),
		Err:         err,
		Info:        infostr,
	}
}

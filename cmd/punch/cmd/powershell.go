package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tomsteele/wholepunch/pkg/wp"
)

var (
	flPowershellHTTPGetSpoofHostHeaderServerPort string
	flPowershellHTTPGetSpoofHostHeaderUserAgent  string
	flPowershellHTTPGetSpoofHostHeaderHostHeader string
	flPowershellHTTPGetSpoofHostHeaderURLPath    string
	flPowershellHTTPGetSpoofHostHeaderUseTLS     bool
)

func init() {
	beaconCmd.AddCommand(powershellHTTPGetSpoofHostHeaderCmd)
	powershellHTTPGetSpoofHostHeaderCmd.PersistentFlags().StringVar(&flPowershellHTTPGetSpoofHostHeaderServerPort, "server-port", "80", "HTTP port to connect to.")
	powershellHTTPGetSpoofHostHeaderCmd.PersistentFlags().StringVar(&flPowershellHTTPGetSpoofHostHeaderHostHeader, "host-header", "www.microsoft.com", "HTTP Host header to use during HTTP request.")
	powershellHTTPGetSpoofHostHeaderCmd.PersistentFlags().StringVar(&flPowershellHTTPGetSpoofHostHeaderUserAgent, "user-agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; AS; rv:11.0) like Gecko", "User-Agent to use during HTTP request.")
	powershellHTTPGetSpoofHostHeaderCmd.PersistentFlags().StringVar(&flPowershellHTTPGetSpoofHostHeaderURLPath, "path", "/", "URL path to use during HTTP request.")
	powershellHTTPGetSpoofHostHeaderCmd.PersistentFlags().BoolVar(&flPowershellHTTPGetSpoofHostHeaderUseTLS, "tls", false, "Use TLS for HTTP connection.")
}

func powershellHTTPGetSpoofHostHeader(cmd *cobra.Command, args []string) {
	opts := wp.BeaconOptions{
		DestinationServerAddress: fmt.Sprintf("%s:%s", flBeaconServerAddr, flPowershellHTTPGetSpoofHostHeaderServerPort),
	}
	var b wp.Beacon
	b = &wp.BeaconPowershellHTTPGetSpoofHostHeader{
		HostHeaderName: flPowershellHTTPGetSpoofHostHeaderHostHeader,
		UserAgent:      flPowershellHTTPGetSpoofHostHeaderUserAgent,
		Path:           flPowershellHTTPGetSpoofHostHeaderURLPath,
	}
	if flPowershellHTTPGetSpoofHostHeaderUseTLS {
		b = &wp.BeaconPowershellHTTPSGetSpoofHostHeader{
			HostHeaderName: flPowershellHTTPGetSpoofHostHeaderHostHeader,
			UserAgent:      flPowershellHTTPGetSpoofHostHeaderUserAgent,
			Path:           flPowershellHTTPGetSpoofHostHeaderURLPath,
		}
	}
	ok, err := wp.RunBeacon(b, &opts)
	result := wp.MakeBeaconResult(ok, err, b)
	wp.WriteTableBeaconResults(os.Stdout, []wp.BeaconResult{result}, flBeaconFilterFalse)
}

var powershellHTTPGetSpoofHostHeaderCmd = &cobra.Command{
	Use:   "powershell-http-get-spoof-host-header",
	Short: "Send an HTTP(S) GET request using Powershell using a modified Host header that does not match the destination.",
	Run:   powershellHTTPGetSpoofHostHeader,
}

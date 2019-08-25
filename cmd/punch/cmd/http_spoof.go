package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tomsteele/wholepunch/pkg/wp"
)

var (
	flHTTPGetSpoofHostHeaderServerPort string
	flHTTPGetSpoofHostHeaderUserAgent  string
	flHTTPGetSpoofHostHeaderHostHeader string
	flHTTPGetSpoofHostHeaderURLPath    string
)

func init() {
	beaconCmd.AddCommand(httpGetSpoofHostHeaderCmd)
	httpGetSpoofHostHeaderCmd.PersistentFlags().StringVar(&flHTTPGetSpoofHostHeaderServerPort, "server-port", "80", "HTTP port to connect to.")
	httpGetSpoofHostHeaderCmd.PersistentFlags().StringVar(&flHTTPGetSpoofHostHeaderHostHeader, "host-header", "www.microsoft.com", "HTTP Host header to use during HTTP request.")
	httpGetSpoofHostHeaderCmd.PersistentFlags().StringVar(&flHTTPGetSpoofHostHeaderUserAgent, "user-agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; AS; rv:11.0) like Gecko", "User-Agent to use during HTTP request.")
	httpGetSpoofHostHeaderCmd.PersistentFlags().StringVar(&flHTTPGetSpoofHostHeaderURLPath, "path", "/", "URL path to use during HTTP request.")
}

func httpGetSpoofHostHeader(cmd *cobra.Command, args []string) {
	opts := wp.BeaconOptions{
		DestinationServerAddress: fmt.Sprintf("%s:%s", flBeaconServerAddr, flHTTPGetSpoofHostHeaderServerPort),
	}
	b := wp.BeaconHTTPGetSpoofHostHeader{
		HostHeaderName: flHTTPGetSpoofHostHeaderHostHeader,
		UserAgent:      flHTTPGetSpoofHostHeaderUserAgent,
		Path:           flHTTPGetSpoofHostHeaderURLPath,
	}
	ok, err := wp.RunBeacon(&b, &opts)
	result := wp.MakeBeaconResult(ok, err, &b)
	wp.WriteTableBeaconResults(os.Stdout, []wp.BeaconResult{result}, flBeaconFilterFalse)
}

var httpGetSpoofHostHeaderCmd = &cobra.Command{
	Use:   "http-get-spoof-host-header",
	Short: "Send an HTTP GET requuest using a modified Host header that does not match the destination.",
	Run:   httpGetSpoofHostHeader,
}

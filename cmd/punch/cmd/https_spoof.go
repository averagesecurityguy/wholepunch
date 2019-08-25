package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tomsteele/wholepunch/pkg/wp"
)

var (
	flHTTPSGetSpoofHostHeaderServerPort string
	flHTTPSGetSpoofHostHeaderUserAgent  string
	flHTTPSGetSpoofHostHeaderHostHeader string
	flHTTPSGetSpoofHostHeaderURLPath    string
)

func init() {
	beaconCmd.AddCommand(httpsGetSpoofHostHeaderCmd)
	httpsGetSpoofHostHeaderCmd.PersistentFlags().StringVar(&flHTTPSGetSpoofHostHeaderServerPort, "server-port", "80", "HTTPS port to connect to.")
	httpsGetSpoofHostHeaderCmd.PersistentFlags().StringVar(&flHTTPSGetSpoofHostHeaderHostHeader, "host-header", "www.microsoft.com", "HTTPS Host header to use during HTTPS request.")
	httpsGetSpoofHostHeaderCmd.PersistentFlags().StringVar(&flHTTPSGetSpoofHostHeaderUserAgent, "user-agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; AS; rv:11.0) like Gecko", "User-Agent to use during HTTPS request.")
	httpsGetSpoofHostHeaderCmd.PersistentFlags().StringVar(&flHTTPSGetSpoofHostHeaderURLPath, "path", "/", "URL path to use during HTTPS request.")
}

func httpsGetSpoofHostHeader(cmd *cobra.Command, args []string) {
	opts := wp.BeaconOptions{
		DestinationServerAddress: fmt.Sprintf("%s:%s", flBeaconServerAddr, flHTTPSGetSpoofHostHeaderServerPort),
	}
	b := wp.BeaconHTTPSGetSpoofHostHeader{
		HostHeaderName: flHTTPSGetSpoofHostHeaderHostHeader,
		UserAgent:      flHTTPSGetSpoofHostHeaderUserAgent,
		Path:           flHTTPSGetSpoofHostHeaderURLPath,
	}
	ok, err := wp.RunBeacon(&b, &opts)
	result := wp.MakeBeaconResult(ok, err, &b)
	wp.WriteTableBeaconResults(os.Stdout, []wp.BeaconResult{result}, flBeaconFilterFalse)
}

var httpsGetSpoofHostHeaderCmd = &cobra.Command{
	Use:   "https-get-spoof-host-header",
	Short: "Send an HTTPS GET requuest using a modified Host header that does not match the destination.",
	Run:   httpsGetSpoofHostHeader,
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tomsteele/wholepunch/pkg/wp"
)

var (
	flHTTPSGetServerPort string
	flHTTPSGetUserAgent  string
	flHTTPSGetURLPath    string
)

func init() {
	beaconCmd.AddCommand(httpsGetCmd)
	httpsGetCmd.PersistentFlags().StringVar(&flHTTPSGetServerPort, "server-port", "80", "HTTP port to connect to.")
	httpsGetCmd.PersistentFlags().StringVar(&flHTTPSGetUserAgent, "user-agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; AS; rv:11.0) like Gecko", "User-Agent to use during HTTP request.")
	httpsGetCmd.PersistentFlags().StringVar(&flHTTPSGetURLPath, "path", "/", "URL path to use during HTTP request.")
}

func httpsGet(cmd *cobra.Command, args []string) {
	opts := wp.BeaconOptions{
		DestinationServerAddress: fmt.Sprintf("%s:%s", flBeaconServerAddr, flHTTPSGetServerPort),
	}
	b := wp.BeaconHTTPSGet{
		UserAgent: flHTTPSGetUserAgent,
		Path:      flHTTPSGetURLPath,
	}
	ok, err := wp.RunBeacon(&b, &opts)
	result := wp.MakeBeaconResult(ok, err, &b)
	wp.WriteTableBeaconResults(os.Stdout, []wp.BeaconResult{result}, flBeaconFilterFalse)
}

var httpsGetCmd = &cobra.Command{
	Use:   "https-get",
	Short: "Send an HTTP GET request using TLS to a desired server.",
	Run:   httpsGet,
}

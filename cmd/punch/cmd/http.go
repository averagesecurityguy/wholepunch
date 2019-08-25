package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tomsteele/wholepunch/pkg/wp"
)

var (
	flHTTPGetServerPort string
	flHTTPGetUserAgent  string
	flHTTPGetURLPath    string
)

func init() {
	beaconCmd.AddCommand(httpGetCmd)
	httpGetCmd.PersistentFlags().StringVar(&flHTTPGetServerPort, "server-port", "80", "HTTP port to connect to.")
	httpGetCmd.PersistentFlags().StringVar(&flHTTPGetUserAgent, "user-agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; AS; rv:11.0) like Gecko", "User-Agent to use during HTTP request.")
	httpGetCmd.PersistentFlags().StringVar(&flHTTPGetURLPath, "path", "/", "URL path to use during HTTP request.")
}

func httpGet(cmd *cobra.Command, args []string) {
	opts := wp.BeaconOptions{
		DestinationServerAddress: fmt.Sprintf("%s:%s", flBeaconServerAddr, flHTTPGetServerPort),
	}
	b := wp.BeaconHTTPGet{
		UserAgent: flHTTPGetUserAgent,
		Path:      flHTTPGetURLPath,
	}
	ok, err := wp.RunBeacon(&b, &opts)
	result := wp.MakeBeaconResult(ok, err, &b)
	wp.WriteTableBeaconResults(os.Stdout, []wp.BeaconResult{result}, flBeaconFilterFalse)
}

var httpGetCmd = &cobra.Command{
	Use:   "http-get",
	Short: "Send an HTTP GET request to a desired server.",
	Run:   httpGet,
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tomsteele/wholepunch/pkg/wp"
)

var (
	flTLSSpoofSNIServerPort string
	flTLSSpoofServerName    string
)

func init() {
	beaconCmd.AddCommand(tlsSpoofSNICmd)
	tlsSpoofSNICmd.PersistentFlags().StringVar(&flTLSSpoofSNIServerPort, "server-port", "443", "Server port to connect to during TLS attempt.")
	tlsSpoofSNICmd.PersistentFlags().StringVar(&flTLSSpoofServerName, "server-name", "www.microsoft.com", "Server Name to use in the SNI client hello.")
}

func tlsSpoofSNI(cmd *cobra.Command, args []string) {
	opts := wp.BeaconOptions{
		DestinationServerAddress: fmt.Sprintf("%s:%s", flBeaconServerAddr, flTLSSpoofSNIServerPort),
	}
	b := wp.BeaconSpoofSNI{
		ServerName: flTLSSpoofServerName,
	}
	ok, err := wp.RunBeacon(&b, &opts)
	result := wp.MakeBeaconResult(ok, err, &b)
	wp.WriteTableBeaconResults(os.Stdout, []wp.BeaconResult{result}, flBeaconFilterFalse)
}

var tlsSpoofSNICmd = &cobra.Command{
	Use:   "tls-spoof-sni",
	Short: "Initiate a TLS request to the destination server",
	Run:   tlsSpoofSNI,
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tomsteele/wholepunch/pkg/wp"
)

var (
	flTLSTCPServerPort string
)

func init() {
	beaconCmd.AddCommand(tlstcpCmd)
	tlstcpCmd.PersistentFlags().StringVar(&flTLSTCPServerPort, "server-port", "443", "Server port to connect to during TLS attempt.")
}

func tlstcp(cmd *cobra.Command, args []string) {
	opts := wp.BeaconOptions{
		DestinationServerAddress: fmt.Sprintf("%s:%s", flBeaconServerAddr, flTLSTCPServerPort),
	}
	b := wp.BeaconTLSConnect{}
	ok, err := wp.RunBeacon(&b, &opts)
	result := wp.MakeBeaconResult(ok, err, &b)
	wp.WriteTableBeaconResults(os.Stdout, []wp.BeaconResult{result}, flBeaconFilterFalse)
}

var tlstcpCmd = &cobra.Command{
	Use:   "tls-tcp",
	Short: "Initiate a TLS/TCP request to the destination server",
	Run:   tlstcp,
}

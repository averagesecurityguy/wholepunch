package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tomsteele/wholepunch/pkg/wp"
	"github.com/tomsteele/xplode"
)

var (
	flTCPServerPort string
	flTCPTimeout    int
)

func init() {
	beaconCmd.AddCommand(tcpCmd)
	tcpCmd.PersistentFlags().StringVar(&flTCPServerPort, "p", "1-1024", "NMap style port string.")
	tcpCmd.PersistentFlags().IntVar(&flTCPTimeout, "timeout", 500, "Timeout in milliseconds.")
}

func tcp(cmd *cobra.Command, args []string) {
	ports, err := xplode.Parse(flTCPServerPort)
	if err != nil {
		fmt.Println("There was an error parsing the port string.")
		fmt.Println(err)
		os.Exit(1)
	}
	results := []wp.BeaconResult{}
	for _, p := range ports {
		b := wp.BeaconTCP{
			Timeout: flTCPTimeout,
		}
		opts := wp.BeaconOptions{
			DestinationServerAddress: fmt.Sprintf("%s:%d", flBeaconServerAddr, p),
		}
		ok, err := wp.RunBeacon(&b, &opts)
		result := wp.MakeBeaconResult(ok, err, &b)
		results = append(results, result)
	}
	wp.WriteTableBeaconResults(os.Stdout, results, flBeaconFilterFalse)
}

var tcpCmd = &cobra.Command{
	Use:   "tcp",
	Short: "Initiates TCP connections to the destination server.",
	Run:   tcp,
}

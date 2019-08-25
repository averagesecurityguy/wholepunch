package cmd

import "github.com/spf13/cobra"

var (
	flBeaconServerAddr       string
	flBeaconFilterFalse      bool
	flBeaconValidateDownload string
	flBeaconUploadData       string
)

func init() {
	rootCmd.AddCommand(beaconCmd)
	beaconCmd.PersistentFlags().StringVar(&flBeaconServerAddr, "server-address", "", "Server address used for direct communications.")
	beaconCmd.PersistentFlags().StringVar(&flBeaconValidateDownload, "data-download-hash", "", "SHA256 of data to download.")
	beaconCmd.PersistentFlags().StringVar(&flBeaconUploadData, "data-upload", "", "Data upload scheme. There is an entire language here.")
	beaconCmd.PersistentFlags().BoolVar(&flBeaconFilterFalse, "filter-false", false, "Do not print results for unsuccessful connections.")
}

func beacon(cmd *cobra.Command, args []string) {
}

var beaconCmd = &cobra.Command{
	Use:   "beacon",
	Short: "Executes an egress assessment from an endpoint",
	Long:  "Executes an egress assessment from an endpoint",
	Run:   beacon,
}

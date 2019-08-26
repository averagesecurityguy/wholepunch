package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tomsteele/wholepunch/pkg/wp"
)

var (
	flMQTTGetServerPort      string
	flMQTTPrivateKeyPath     string
	flMQTTCertificatePath    string
	flMQTTClientID           string
	flMQTTInsecureSkipVerify bool
	flMQTTCACertificatePath  string
)

func init() {
	beaconCmd.AddCommand(mqttCmd)
	mqttCmd.PersistentFlags().StringVar(&flMQTTGetServerPort, "server-port", "443", "MQTT port to connect to.")
	mqttCmd.PersistentFlags().StringVar(&flMQTTPrivateKeyPath, "key", "", "TLS client private key.")
	mqttCmd.PersistentFlags().StringVar(&flMQTTCertificatePath, "cert", "", "TLS client certificate path.")
	mqttCmd.PersistentFlags().StringVar(&flMQTTClientID, "id", "", "MQTT client id.")
	mqttCmd.PersistentFlags().StringVar(&flMQTTCACertificatePath, "ca", "", "TLS root CA.")
	mqttCmd.PersistentFlags().BoolVar(&flMQTTInsecureSkipVerify, "insecure", false, "TLS insecure skip verify.")
}

func mqtt(cmd *cobra.Command, args []string) {
	opts := wp.BeaconOptions{
		DestinationServerAddress: fmt.Sprintf("%s:%s", flBeaconServerAddr, flMQTTGetServerPort),
	}
	b := wp.BeaconMQTT{
		PrivateKeyPath:     flMQTTPrivateKeyPath,
		CertificatePath:    flMQTTCertificatePath,
		InsecureSkipVerify: flMQTTInsecureSkipVerify,
		ClientID:           flMQTTClientID,
		CACertificatePath:  flMQTTCACertificatePath,
	}
	ok, err := wp.RunBeacon(&b, &opts)
	result := wp.MakeBeaconResult(ok, err, &b)
	wp.WriteTableBeaconResults(os.Stdout, []wp.BeaconResult{result}, flBeaconFilterFalse)
}

var mqttCmd = &cobra.Command{
	Use:   "mqtt",
	Short: "Connect to a destination server over MQTT.",
	Run:   mqtt,
}

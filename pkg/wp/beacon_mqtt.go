package wp

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// BeaconMQTT connects to an MQTT server using TLS.
type BeaconMQTT struct {
	PrivateKeyPath     string
	CertificatePath    string
	CACertificatePath  string
	ClientID           string
	InsecureSkipVerify bool
	ServerAddr         string
}

// Name returns the name of the module.
func (b *BeaconMQTT) Name() string {
	return "mqtt"
}

// Destination returns the server that was connected to
func (b *BeaconMQTT) Destination() string {
	return b.ServerAddr
}

// Success returns a formatted string indicating a successfull connection.
func (b *BeaconMQTT) Success() string {
	return fmt.Sprintf("The agent was allowed to communicate with %s over MQTT", b.ServerAddr)
}

// Setup is used to initilize instance variables from BeaconOptions.
func (b *BeaconMQTT) Setup(o *BeaconOptions) error {
	b.ServerAddr = o.DestinationServerAddress
	return nil
}

// Send connects to an MQTT server.
func (b *BeaconMQTT) Send() (bool, error) {
	tlsCert, err := tls.LoadX509KeyPair(b.CertificatePath, b.PrivateKeyPath)
	if err != nil {
		return false, err
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
	}
	if b.CACertificatePath != "" {
		certs := x509.NewCertPool()
		caPem, err := ioutil.ReadFile(b.CACertificatePath)
		if err != nil {
			return false, err
		}
		certs.AppendCertsFromPEM(caPem)
	}
	if b.InsecureSkipVerify {
		tlsConfig.InsecureSkipVerify = true
	}
	tlsConfig.ServerName = "a74rjmjw9m32l-ats.iot.us-west-2.amazonaws.com"
	mqtturl := fmt.Sprintf("ssl://%s", b.ServerAddr)
	mqttOpts := mqtt.NewClientOptions()
	mqttOpts.AddBroker(mqtturl)
	mqttOpts.SetMaxReconnectInterval(1 * time.Second)
	mqttOpts.SetClientID(b.ClientID)
	mqttOpts.SetTLSConfig(tlsConfig)
	c := mqtt.NewClient(mqttOpts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		return false, token.Error()
	}
	c.Disconnect(0)
	return true, nil
}

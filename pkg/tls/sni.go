package tls

import (
	"crypto/tls"

	log "github.com/sirupsen/logrus"
)

// SendSpoofSNI creates a TLS connection with a ServerName in the Client Hello portion
// of the TLS handshake set to your provided hostname.
func SendSpoofSNI(serverName, serverAddr string) (bool, error) {
	config := tls.Config{
		ServerName:         serverName,
		InsecureSkipVerify: true,
	}
	conn, err := tls.Dial("tcp", serverAddr, &config)
	if err != nil {
		log.Debugf("error during tls dial %s", err.Error())
		return false, err
	}
	if err := conn.Close(); err != nil {
		log.Debugf("error closing tls conn %s", err.Error())
	}
	return true, nil
}

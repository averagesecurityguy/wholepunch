package tls

import (
	"crypto/tls"

	log "github.com/sirupsen/logrus"
)

// DialInsecureTCP connects to a TLS server using insecure-skip-verify.
func DialInsecureTCP(serverAddr string) (bool, error) {
	conn, err := tls.Dial("tcp", serverAddr, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Debugf("error during tls dial %s", err.Error())
		return false, err
	}
	if err := conn.Close(); err != nil {
		log.Debugf("error closing tls conn %s", err.Error())
	}
	return true, nil
}

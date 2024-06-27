package ntp

import (
	"github.com/beevik/ntp"
	log "github.com/sirupsen/logrus"
)

// DialInsecureTCP connects to a TLS server using insecure-skip-verify.
func GetCurrentTime(serverAddr string) (bool, error) {
	_, err := ntp.Query(serverAddr)
	if err != nil {
		log.Debugf("error during ntp query %s", err.Error())
		return false, err
	}

	return true, nil
}

package wp

import (
	"net/url"
	"strconv"
	"strings"
)

func tcpBeaconLayerFromURL(uri string) BeaconLayerData {
	b := BeaconLayerData{
		NetworkLayer:   "IP",
		TransportLayer: "TCP",
	}
	u, err := url.Parse(uri)
	if err != nil {
		return b
	}
	b.AppLayer = strings.ToUpper(u.Scheme)
	b.TransportLayerPort, _ = strconv.Atoi(u.Port())
	return b
}

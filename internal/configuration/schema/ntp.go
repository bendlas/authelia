package schema

import (
	"net/url"
	"time"
)

// NTPConfiguration represents the configuration related to ntp server.
type NTPConfiguration struct {
	Address             *AddressUDP   `koanf:"address" json:"address"`
	Version             int           `koanf:"version" json:"version"`
	MaximumDesync       time.Duration `koanf:"max_desync" json:"max_desync"`
	DisableStartupCheck bool          `koanf:"disable_startup_check" json:"disable_startup_check"`
	DisableFailure      bool          `koanf:"disable_failure" json:"disable_failure"`
}

// DefaultNTPConfiguration represents default configuration parameters for the NTP server.
var DefaultNTPConfiguration = NTPConfiguration{
	Address:       &AddressUDP{Address{valid: true, socket: false, port: 123, url: &url.URL{Scheme: AddressSchemeUDP, Host: "time.cloudflare.com:123"}}},
	Version:       4,
	MaximumDesync: time.Second * 3,
}

package schema

import (
	"time"
)

// TLSConfig is a representation of the TLS configuration.
type TLSConfig struct {
	MinimumVersion TLSVersion `koanf:"minimum_version" json:"minimum_version"`
	MaximumVersion TLSVersion `koanf:"maximum_version" json:"maximum_version"`

	SkipVerify bool   `koanf:"skip_verify" json:"skip_verify"`
	ServerName string `koanf:"server_name" json:"server_name"`

	PrivateKey       CryptographicPrivateKey `koanf:"private_key" json:"private_key"`
	CertificateChain X509CertificateChain    `koanf:"certificate_chain" json:"certificate_chain"`
}

// TLSCertificateConfig is a representation of the TLS Certificate configuration.
type TLSCertificateConfig struct {
	Key              CryptographicPrivateKey `koanf:"key" json:"key"`
	CertificateChain X509CertificateChain    `koanf:"certificate_chain" json:"certificate_chain"`
}

// ServerTimeouts represents server timeout configurations.
type ServerTimeouts struct {
	Read  time.Duration `koanf:"read" json:"read"`
	Write time.Duration `koanf:"write" json:"write"`
	Idle  time.Duration `koanf:"idle" json:"idle"`
}

// ServerBuffers represents server buffer configurations.
type ServerBuffers struct {
	Read  int `koanf:"read" json:"read"`
	Write int `koanf:"write" json:"write"`
}

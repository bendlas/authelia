package schema

// TOTPConfiguration represents the configuration related to TOTP options.
type TOTPConfiguration struct {
	Disable    bool   `koanf:"disable" json:"disable"`
	Issuer     string `koanf:"issuer" json:"issuer"`
	Algorithm  string `koanf:"algorithm" json:"algorithm"`
	Digits     uint   `koanf:"digits" json:"digits"`
	Period     uint   `koanf:"period" json:"period"`
	Skew       *uint  `koanf:"skew" json:"skew"`
	SecretSize uint   `koanf:"secret_size" json:"secret_size"`
}

var defaultOtpSkew = uint(1)

// DefaultTOTPConfiguration represents default configuration parameters for TOTP generation.
var DefaultTOTPConfiguration = TOTPConfiguration{
	Issuer:     "Authelia",
	Algorithm:  TOTPAlgorithmSHA1,
	Digits:     6,
	Period:     30,
	Skew:       &defaultOtpSkew,
	SecretSize: TOTPSecretSizeDefault,
}

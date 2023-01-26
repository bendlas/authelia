package schema

// DuoAPIConfiguration represents the configuration related to Duo API.
type DuoAPIConfiguration struct {
	Disable              bool   `koanf:"disable" json:"disable"`
	Hostname             string `koanf:"hostname" json:"hostname"`
	IntegrationKey       string `koanf:"integration_key" json:"integration_key"`
	SecretKey            string `koanf:"secret_key" json:"secret_key"`
	EnableSelfEnrollment bool   `koanf:"enable_self_enrollment" json:"enable_self_enrollment"`
}

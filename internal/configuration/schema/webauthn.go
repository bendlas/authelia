package schema

import (
	"time"

	"github.com/go-webauthn/webauthn/protocol"
)

// WebAuthnConfiguration represents the webauthn config.
type WebAuthnConfiguration struct {
	Disable     bool   `koanf:"disable" json:"disable"`
	DisplayName string `koanf:"display_name" json:"display_name"`

	ConveyancePreference protocol.ConveyancePreference        `koanf:"attestation_conveyance_preference" json:"attestation_conveyance_preference"`
	UserVerification     protocol.UserVerificationRequirement `koanf:"user_verification" json:"user_verification"`

	Timeout time.Duration `koanf:"timeout" json:"timeout"`
}

// DefaultWebAuthnConfiguration describes the default values for the WebAuthnConfiguration.
var DefaultWebAuthnConfiguration = WebAuthnConfiguration{
	DisplayName: "Authelia",
	Timeout:     time.Second * 60,

	ConveyancePreference: protocol.PreferIndirectAttestation,
	UserVerification:     protocol.VerificationPreferred,
}

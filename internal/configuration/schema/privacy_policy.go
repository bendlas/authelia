package schema

import (
	"net/url"
)

// PrivacyPolicy is the privacy policy configuration.
type PrivacyPolicy struct {
	Enabled               bool     `koanf:"enabled" json:"enabled"`
	RequireUserAcceptance bool     `koanf:"require_user_acceptance" json:"require_user_acceptance"`
	PolicyURL             *url.URL `koanf:"policy_url" json:"policy_url"`
}

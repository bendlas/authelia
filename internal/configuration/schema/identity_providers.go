package schema

import (
	"crypto/rsa"
	"net/url"
	"time"
)

// IdentityProvidersConfiguration represents the IdentityProviders 2.0 configuration for Authelia.
type IdentityProvidersConfiguration struct {
	OIDC *OpenIDConnectConfiguration `koanf:"oidc" json:"oidc"`
}

// OpenIDConnectConfiguration configuration for OpenID Connect.
type OpenIDConnectConfiguration struct {
	HMACSecret             string               `koanf:"hmac_secret" json:"hmac_secret"`
	IssuerCertificateChain X509CertificateChain `koanf:"issuer_certificate_chain" json:"issuer_certificate_chain"`
	IssuerPrivateKey       *rsa.PrivateKey      `koanf:"issuer_private_key" json:"issuer_private_key"`

	AccessTokenLifespan   time.Duration `koanf:"access_token_lifespan" json:"access_token_lifespan"`
	AuthorizeCodeLifespan time.Duration `koanf:"authorize_code_lifespan" json:"authorize_code_lifespan"`
	IDTokenLifespan       time.Duration `koanf:"id_token_lifespan" json:"id_token_lifespan"`
	RefreshTokenLifespan  time.Duration `koanf:"refresh_token_lifespan" json:"refresh_token_lifespan"`

	EnableClientDebugMessages bool `koanf:"enable_client_debug_messages" json:"enable_client_debug_messages"`
	MinimumParameterEntropy   int  `koanf:"minimum_parameter_entropy" json:"minimum_parameter_entropy"`

	EnforcePKCE              string `koanf:"enforce_pkce" json:"enforce_pkce"`
	EnablePKCEPlainChallenge bool   `koanf:"enable_pkce_plain_challenge" json:"enable_pkce_plain_challenge"`

	CORS OpenIDConnectCORSConfiguration `koanf:"cors" json:"cors"`
	PAR  OpenIDConnectPARConfiguration  `koanf:"pushed_authorizations" json:"pushed_authorizations"`

	Clients []OpenIDConnectClientConfiguration `koanf:"clients" json:"clients"`
}

// OpenIDConnectPARConfiguration represents an OpenID Connect PAR config.
type OpenIDConnectPARConfiguration struct {
	Enforce         bool          `koanf:"enforce"`
	ContextLifespan time.Duration `koanf:"context_lifespan"`
}

// OpenIDConnectCORSConfiguration represents an OpenID Connect CORS config.
type OpenIDConnectCORSConfiguration struct {
	Endpoints      []string  `koanf:"endpoints" json:"endpoints"`
	AllowedOrigins []url.URL `koanf:"allowed_origins" json:"allowed_origins"`

	AllowedOriginsFromClientRedirectURIs bool `koanf:"allowed_origins_from_client_redirect_uris" json:"allowed_origins_from_client_redirect_uris"`
}

// OpenIDConnectClientConfiguration configuration for an OpenID Connect client.
type OpenIDConnectClientConfiguration struct {
	ID               string          `koanf:"id" json:"id"`
	Description      string          `koanf:"description" json:"description"`
	Secret           *PasswordDigest `koanf:"secret" json:"secret"`
	SectorIdentifier url.URL         `koanf:"sector_identifier" json:"sector_identifier"`
	Public           bool            `koanf:"public" json:"public"`

	RedirectURIs []string `koanf:"redirect_uris" json:"redirect_uris"`

	Audience      []string `koanf:"audience" json:"audience"`
	Scopes        []string `koanf:"scopes" json:"scopes"`
	GrantTypes    []string `koanf:"grant_types" json:"grant_types"`
	ResponseTypes []string `koanf:"response_types" json:"response_types"`
	ResponseModes []string `koanf:"response_modes" json:"response_modes"`

	TokenEndpointAuthMethod string `koanf:"token_endpoint_auth_method" json:"token_endpoint_auth_method"`

	Policy string `koanf:"authorization_policy" json:"policy"`

	EnforcePAR  bool `koanf:"enforce_par" json:"enforce_par"`
	EnforcePKCE bool `koanf:"enforce_pkce" json:"enforce_pkce"`

	PKCEChallengeMethod      string `koanf:"pkce_challenge_method" json:"pkce_challenge_method"`
	UserinfoSigningAlgorithm string `koanf:"userinfo_signing_algorithm" json:"userinfo_signing_algorithm"`

	ConsentMode                  string         `koanf:"consent_mode" json:"consent_mode"`
	ConsentPreConfiguredDuration *time.Duration `koanf:"pre_configured_consent_duration" json:"pre_configured_consent_duration"`
}

// DefaultOpenIDConnectConfiguration contains defaults for OIDC.
var DefaultOpenIDConnectConfiguration = OpenIDConnectConfiguration{
	AccessTokenLifespan:   time.Hour,
	AuthorizeCodeLifespan: time.Minute,
	IDTokenLifespan:       time.Hour,
	RefreshTokenLifespan:  time.Minute * 90,
	EnforcePKCE:           "public_clients_only",
}

var defaultOIDCClientConsentPreConfiguredDuration = time.Hour * 24 * 7

// DefaultOpenIDConnectClientConfiguration contains defaults for OIDC Clients.
var DefaultOpenIDConnectClientConfiguration = OpenIDConnectClientConfiguration{
	Policy:        "two_factor",
	Scopes:        []string{"openid", "groups", "profile", "email"},
	ResponseTypes: []string{"code"},
	ResponseModes: []string{"form_post"},

	UserinfoSigningAlgorithm:     "none",
	ConsentMode:                  "auto",
	ConsentPreConfiguredDuration: &defaultOIDCClientConsentPreConfiguredDuration,
}

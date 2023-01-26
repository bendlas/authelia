package schema

import (
	"net/url"
	"time"
)

// ServerConfiguration represents the configuration of the http server.
type ServerConfiguration struct {
	Address            *AddressTCP `koanf:"address"`
	Path               string `koanf:"path" json:"path"`
	AssetPath          string `koanf:"asset_path" json:"asset_path"`
	DisableHealthcheck bool   `koanf:"disable_healthcheck" json:"disable_healthcheck"`

	TLS       ServerTLS       `koanf:"tls" json:"tls"`
	Headers   ServerHeaders   `koanf:"headers" json:"headers"`
	Endpoints ServerEndpoints `koanf:"endpoints" json:"endpoints"`

	Buffers  ServerBuffers  `koanf:"buffers" json:"buffers"`
	Timeouts ServerTimeouts `koanf:"timeouts" json:"timeouts"`

	// Deprecated: use address instead.
	Host string `koanf:"host"`

	// Deprecated: use address instead.
	Port int `koanf:"port"`
}

// ServerEndpoints is the endpoints configuration for the HTTP server.
type ServerEndpoints struct {
	EnablePprof   bool `koanf:"enable_pprof" json:"enable_pprof"`
	EnableExpvars bool `koanf:"enable_expvars" json:"enable_expvars"`

	Authz map[string]ServerAuthzEndpoint `koanf:"authz" json:"authz"`
}

// ServerAuthzEndpoint is the Authz endpoints configuration for the HTTP server.
type ServerAuthzEndpoint struct {
	Implementation string `koanf:"implementation" json:"implementation"`

	AuthnStrategies []ServerAuthzEndpointAuthnStrategy `koanf:"authn_strategies" json:"authn_strategies"`
}

// ServerAuthzEndpointAuthnStrategy is the Authz endpoints configuration for the HTTP server.
type ServerAuthzEndpointAuthnStrategy struct {
	Name string `koanf:"name" json:"name"`
}

// ServerTLS represents the configuration of the http servers TLS options.
type ServerTLS struct {
	Certificate        string   `koanf:"certificate" json:"certificate"`
	Key                string   `koanf:"key" json:"key"`
	ClientCertificates []string `koanf:"client_certificates" json:"client_certificates"`
}

// ServerHeaders represents the customization of the http server headers.
type ServerHeaders struct {
	CSPTemplate string `koanf:"csp_template" json:"csp_template"`
}

// DefaultServerConfiguration represents the default values of the ServerConfiguration.
var DefaultServerConfiguration = ServerConfiguration{
	Address: &AddressTCP{Address{true, false, -1, 9091, &url.URL{Scheme: AddressSchemeTCP, Host: ":9091"}}},
	Buffers: ServerBuffers{
		Read:  4096,
		Write: 4096,
	},
	Timeouts: ServerTimeouts{
		Read:  time.Second * 6,
		Write: time.Second * 6,
		Idle:  time.Second * 30,
	},
	Endpoints: ServerEndpoints{
		Authz: map[string]ServerAuthzEndpoint{
			"legacy": {
				Implementation: "Legacy",
			},
			"auth-request": {
				Implementation: "AuthRequest",
				AuthnStrategies: []ServerAuthzEndpointAuthnStrategy{
					{
						Name: "HeaderAuthRequestProxyAuthorization",
					},
					{
						Name: "CookieSession",
					},
				},
			},
			"forward-auth": {
				Implementation: "ForwardAuth",
				AuthnStrategies: []ServerAuthzEndpointAuthnStrategy{
					{
						Name: "HeaderProxyAuthorization",
					},
					{
						Name: "CookieSession",
					},
				},
			},
			"ext-authz": {
				Implementation: "ExtAuthz",
				AuthnStrategies: []ServerAuthzEndpointAuthnStrategy{
					{
						Name: "HeaderProxyAuthorization",
					},
					{
						Name: "CookieSession",
					},
				},
			},
		},
	},
}

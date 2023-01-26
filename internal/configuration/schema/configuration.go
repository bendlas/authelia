package schema

// Configuration object extracted from YAML configuration file.
type Configuration struct {
	Theme                 string `koanf:"theme" json:"theme" jsonschema:"title=Theme Name,description=The name of the theme to apply to the web UI,default=light,enum=auto,enum=light,enum=dark,enum=grey"`
	CertificatesDirectory string `koanf:"certificates_directory" json:"certificates_directory" jsonschema:"title=Certificates Directory Path,description=The path to a directory which is used to determine the certificates that are trusted"`
	JWTSecret             string `koanf:"jwt_secret" json:"jwt_secret" jsonschema:"title=Secret Key for JWT's,description=Used for signing HS256 JWT's for identity verification,required"`
	DefaultRedirectionURL string `koanf:"default_redirection_url" json:"default_redirection_url" jsonschema:"title=The default redirection URL,description=Used to redirect users when they visit the portal directly"`
	Default2FAMethod      string `koanf:"default_2fa_method" json:"default_2fa_method" jsonschema:"title=Default 2FA method,description=When a user logs in for the first time this is the 2FA method configured for them"`

	Log                   LogConfiguration               `koanf:"log" json:"log"`
	IdentityProviders     IdentityProvidersConfiguration `koanf:"identity_providers" json:"identity_providers"`
	AuthenticationBackend AuthenticationBackend          `koanf:"authentication_backend" json:"authentication_backend"`
	Session               SessionConfiguration           `koanf:"session" json:"session"`
	TOTP                  TOTPConfiguration              `koanf:"totp" json:"totp"`
	DuoAPI                DuoAPIConfiguration            `koanf:"duo_api" json:"duo_api"`
	AccessControl         AccessControlConfiguration     `koanf:"access_control" json:"access_control"`
	NTP                   NTPConfiguration               `koanf:"ntp" json:"ntp"`
	Regulation            RegulationConfiguration        `koanf:"regulation" json:"regulation"`
	Storage               StorageConfiguration           `koanf:"storage" json:"storage"`
	Notifier              NotifierConfiguration          `koanf:"notifier" json:"notifier"`
	Server                ServerConfiguration            `koanf:"server" json:"server"`
	Telemetry             TelemetryConfig                `koanf:"telemetry" json:"telemetry"`
	WebAuthn              WebAuthnConfiguration          `koanf:"webauthn" json:"webauthn"`
	PasswordPolicy        PasswordPolicyConfiguration    `koanf:"password_policy" json:"password_policy"`
	PrivacyPolicy         PrivacyPolicy                  `koanf:"privacy_policy" json:"privacy_policy"`
}

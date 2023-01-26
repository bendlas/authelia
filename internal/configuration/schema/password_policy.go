package schema

// PasswordPolicyStandardParams represents the configuration related to standard parameters of password policy.
type PasswordPolicyStandardParams struct {
	Enabled          bool `koanf:"enabled" json:"enabled"`
	MinLength        int  `koanf:"min_length" json:"min_length"`
	MaxLength        int  `koanf:"max_length" json:"max_length"`
	RequireUppercase bool `koanf:"require_uppercase" json:"require_uppercase"`
	RequireLowercase bool `koanf:"require_lowercase" json:"require_lowercase"`
	RequireNumber    bool `koanf:"require_number" json:"require_number"`
	RequireSpecial   bool `koanf:"require_special" json:"require_special"`
}

// PasswordPolicyZXCVBNParams represents the configuration related to ZXCVBN parameters of password policy.
type PasswordPolicyZXCVBNParams struct {
	Enabled  bool `koanf:"enabled" json:"enabled"`
	MinScore int  `koanf:"min_score" json:"min_score"`
}

// PasswordPolicyConfiguration represents the configuration related to password policy.
type PasswordPolicyConfiguration struct {
	Standard PasswordPolicyStandardParams `koanf:"standard" json:"standard"`
	ZXCVBN   PasswordPolicyZXCVBNParams   `koanf:"zxcvbn" json:"zxcvbn"`
}

// DefaultPasswordPolicyConfiguration is the default password policy configuration.
var DefaultPasswordPolicyConfiguration = PasswordPolicyConfiguration{
	Standard: PasswordPolicyStandardParams{
		Enabled:   false,
		MinLength: 8,
		MaxLength: 0,
	},
	ZXCVBN: PasswordPolicyZXCVBNParams{
		Enabled:  false,
		MinScore: 3,
	},
}

package schema

import (
	"regexp"
)

// AccessControlConfiguration represents the configuration related to ACLs.
type AccessControlConfiguration struct {
	// The default policy if no other policy matches the request.
	DefaultPolicy string `koanf:"default_policy" json:"default_policy" jsonschema:"default=deny,enum=deny,enum=one_factor,enum=two_factor"`

	// Represents a list of named network groups.
	Networks []ACLNetwork `koanf:"networks" json:"networks" jsonschema:"title=Rules List,description=The list of ACL rules to enumerate for requests"`

	// The ACL rules list.
	Rules []ACLRule `koanf:"rules" json:"rules" jsonschema:"title=Rules List,description=The list of ACL rules to enumerate for requests"`
}

// ACLNetwork represents one ACL network group entry.
type ACLNetwork struct {
	Name     string   `koanf:"name" json:"name" jsonschema:"title=Network Name,description=The name of this network to be used in the networks section of the rules section,required"`
	Networks []string `koanf:"networks" json:"networks" jsonschema:"title=Remote IP's or Networks,description=The remote IP's or networks that this rule applies to,required"`
}

// ACLRule represents one ACL rule entry.
type ACLRule struct {
	Domains      []string         `koanf:"domain" json:"domain" jsonschema:"title=Domain Literals,description=The literal domains to match the domain against that this rule applies to"`
	DomainsRegex []regexp.Regexp  `koanf:"domain_regex" json:"domain_regex" jsonschema:"title=Domain Regex Patterns,description=The regex patterns to match the domain against that this rule applies to"`
	Policy       string           `koanf:"policy" json:"policy" jsonschema:"required,enum=bypass,enum=deny,enum=one_factor,enum=two_factor,title=Rule Policy,description=The policy this rule applies when all criteria match"`
	Subjects     [][]string       `koanf:"subject" json:"subject" jsonschema:"title=Subjects,description=The users or groups that this rule applies to"`
	Networks     []string         `koanf:"networks" json:"networks" jsonschema:"title=Remote IP's or Networks,description=The remote IP's or networks that this rule applies to"`
	Resources    []regexp.Regexp  `koanf:"resources" json:"resources" jsonschema:"title=Resources or Paths,description=The resource paths this rule applies to"`
	Methods      []string         `koanf:"methods" json:"methods" jsonschema:"enum=GET,enum=HEAD,enum=POST,enum=PUT,enum=DELETE,enum=CONNECT,enum=OPTIONS,enum=TRACE,enum=PATCH,enum=PROPFIND,enum=PROPPATCH,enum=MKCOL,enum=COPY,enum=MOVE,enum=LOCK,enum=UNLOCK"`
	Query        [][]ACLQueryRule `koanf:"query" json:"query" jsonschema:"title=Query Rules,description=The list of query parameter rules this rule applies to"`
}

// ACLQueryRule represents the ACL query criteria.
type ACLQueryRule struct {
	Operator string `koanf:"operator" json:"operator" jsonschema:"enum=equal,enum=not equal,enum=present,enum=absent,enum=pattern,enum=not pattern"`
	Key      string `koanf:"key" json:"key"`
	Value    any    `koanf:"value" json:"value"`
}

// DefaultACLNetwork represents the default configuration related to access control network group configuration.
var DefaultACLNetwork = []ACLNetwork{
	{
		Name:     "localhost",
		Networks: []string{"127.0.0.1"},
	},
	{
		Name:     "internal",
		Networks: []string{"10.0.0.0/8"},
	},
}

// DefaultACLRule represents the default configuration related to access control rule configuration.
var DefaultACLRule = []ACLRule{
	{
		Domains: []string{"public.example.com"},
		Policy:  "bypass",
	},
	{
		Domains: []string{"singlefactor.example.com"},
		Policy:  "one_factor",
	},
	{
		Domains: []string{"secure.example.com"},
		Policy:  "two_factor",
	},
}

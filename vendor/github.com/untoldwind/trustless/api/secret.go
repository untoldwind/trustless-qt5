package api

// SecretType is the type of the secret to store
type SecretType string

const (
	// SecretTypeLogin is a generic login (of a webpage)
	SecretTypeLogin SecretType = "login"
	// SecretTypeNote is a general secret note
	SecretTypeNote SecretType = "note"
	// SecretTypeLicence is some kind of software licence (key)
	SecretTypeLicence SecretType = "licence"
	// SecretTypePassword is some kind of generic password/key
	SecretTypePassword SecretType = "password"
)

// SecretCurrent contains the current (head) version of a secret
type SecretCurrent struct {
	ID      string         `json:"id"`
	Type    SecretType     `json:"type"`
	Current *SecretVersion `json:"current,omitempty"`
}

// Secret holds all information of a secret (including all previous versions)
type Secret struct {
	SecretCurrent
	Versions SecretVersions `json:"versions"`
}

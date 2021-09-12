// Package jwt provides functions for the signing and parsing of JSON Web Tokens.
package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/BranDebs/Avocado-Backend/internal/secrets"
)

const (
	// ContextKey stores the context key to obtain JWT string within the context.
	ContextKey = "ctx_key_jwt"

	// EmptySignedToken represents an empty token.
	EmptySignedToken = ""

	issuer = "avocadoro"
)

var (
	// ErrSigningMethod is raised if the incoming token is not signed using an expected signing function.
	ErrSigningMethod = errors.New("unexpected signing method")
	// ErrInvalidClaims is raised if claims parsed is not in expected format.
	ErrInvalidClaims = errors.New("unexpected claims")
	// ErrAudienceClaim is raised if the audience is not a valid.
	ErrAudienceClaim = errors.New("unexpected audience claim")
)

// Settings stores JWT settings present in the config file.
type Settings struct {
	SigningKeyFile string `mapstructure:"signing_key_file"`
	TTL            int64  `mapstructure:"ttl"`

	SigningKey []byte
}

// Init initialises the settings.
func (j *Settings) Init() {
	signingKey, _ := secrets.SingleLineKey(j.SigningKeyFile)
	j.SigningKey = signingKey
}

// JWT represents the JWT to be used.
type JWT jwt.StandardClaims

// New initialises a new JWT using email as the subject.
func New(email string, ttlSeconds int64) *JWT {
	ts := time.Now().Unix()
	exp := ts + ttlSeconds

	jwt := JWT{
		Audience:  issuer,
		ExpiresAt: exp,
		IssuedAt:  ts,
		Issuer:    issuer,
		NotBefore: ts,
		Subject:   email,
	}

	return &jwt
}

// Token returns a JWT string to be returned to users.
func (t *JWT) Token(signingKey []byte) string {
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims(*t))

	ss, err := tok.SignedString(signingKey)
	if err != nil {
		return EmptySignedToken
	}

	return ss
}

// Verify verifies the JWT string and returns true if it is valid.
func Verify(settings *Settings, tokenStr string) (bool, error) {
	token, err := jwt.Parse(tokenStr, keyFunc(settings))
	if err != nil {
		return false, fmt.Errorf("failed to parse token string err: %w", err)
	}

	claims, ok := token.Claims.(jwt.StandardClaims)
	if !ok {
		return false, ErrInvalidClaims
	}

	// CVE-2020-26160
	// https://github.com/advisories/GHSA-w73w-5m7g-f7qc
	if claims.Audience == "" {
		return false, ErrAudienceClaim
	}

	return token.Valid, nil
}

func keyFunc(settings *Settings) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrSigningMethod
		}

		return settings.SigningKey, nil
	}
}

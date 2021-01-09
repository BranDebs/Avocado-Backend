package account

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	EMPTY_SIGNED_TOKEN = ""

	issuer = "avocadoro"
)

type JWT jwt.StandardClaims

func NewJWT(email string, ttlSeconds int64) *JWT {
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

func (t *JWT) Token(signingKey []byte) string {
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims(*t))

	ss, err := tok.SignedString(signingKey)
	if err != nil {
		return EMPTY_SIGNED_TOKEN
	}

	return ss
}

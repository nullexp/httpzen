package misc

import (
	"time"
)

type Claims interface {
	Valid() error
}

type JwtClaim interface {
	GetExpireTime() int64
	GetSubject() string
	GetIssuer() string
	GetAudience() []string
	GetIssuedAt() int64
	GetIdentity() string
	IsExpired() bool
}

func NewValidJwtClaim() JwtClaim {
	return StandardClaims{ExpiresAt: time.Now().Add(10 * time.Minute).Unix()}
}

type StandardClaims struct {
	Audience  []string `json:"aud,omitempty"`
	ExpiresAt int64    `json:"exp,omitempty"`
	Identity  string   `json:"jti,omitempty"`
	IssuedAt  int64    `json:"iat,omitempty"`
	Issuer    string   `json:"iss,omitempty"`
	NotBefore int64    `json:"nbf,omitempty"`
	Subject   string   `json:"sub,omitempty"`
}

func (c StandardClaims) Valid() error {
	return nil
}

func (c StandardClaims) GetExpireTime() int64 {
	return c.ExpiresAt
}

func (c StandardClaims) GetSubject() string {
	return c.Subject
}

func (c StandardClaims) GetIssuer() string {
	return c.Issuer
}

func (c StandardClaims) GetAudience() []string {
	return c.Audience
}

func (c StandardClaims) GetIssuedAt() int64 {
	return c.IssuedAt
}

func (c StandardClaims) GetIdentity() string {
	return c.Identity
}

func (c StandardClaims) IsExpired() bool {
	return time.Now().Unix() > c.ExpiresAt
}

package protocol

import "github.com/nullexp/httpzen/misc"

type Authenticator interface {
	GetModel(token string) (misc.JwtClaim, error)
	CheckToken(token string) (bool, error)
}

type testAuthenticator struct {
	claim  misc.JwtClaim
	result bool
}

func (ta testAuthenticator) GetModel(token string) (misc.JwtClaim, error) {
	return ta.claim, nil
}

func (ta testAuthenticator) CheckToken(token string) (bool, error) {
	return ta.result, nil
}

func NewTestAuthenticator(claim misc.JwtClaim, rs bool) testAuthenticator {
	return testAuthenticator{claim: claim, result: rs}
}

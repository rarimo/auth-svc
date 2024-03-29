package auth

import (
	"github.com/rarimo/auth-svc/resources"
)

func UserGrant(userDID string) Grant {
	return func(claim resources.Claim) bool {
		return claim.User == userDID
	}
}

package types

import jwt "github.com/dgrijalva/jwt-go"

// MidgardClaims is a struct to represeting the data some a Midgard issued claims.
type MidgardClaims struct {
	Realm  string            `msgpack:"realm" json:"realm"`
	Quota  int               `msgpack:"quota,omitempty" json:"quota,omitempty"`
	Data   map[string]string `msgpack:"data" json:"data"`
	Opaque map[string]string `msgpack:"opaque,omitempty" json:"opaque,omitempty"`

	jwt.StandardClaims
}

// NewMidgardClaims returns a new Claims.
func NewMidgardClaims() *MidgardClaims {
	return &MidgardClaims{
		Data:           map[string]string{},
		StandardClaims: jwt.StandardClaims{},
	}
}

// ServiceToken is a struct to represent the service tokens issued by the system.
type ServiceToken struct {
	User    map[string]interface{} `json:"user,omitempty"`
	Service map[string]interface{} `json:"service,omitempty"`

	jwt.StandardClaims
}

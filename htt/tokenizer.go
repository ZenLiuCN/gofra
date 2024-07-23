package htt

import (
	"github.com/ZenLiuCN/gofra/conf"
	"github.com/golang-jwt/jwt/v5"
)

var (
	method jwt.SigningMethod
)

// JwtInitialize initialize jwt
func JwtInitialize() {
	ConfigJwt(conf.GetConfig())
}

// ConfigJwt manually config jwt token.
func ConfigJwt(conf conf.Config) {
	method = jwt.GetSigningMethod(conf.GetString("jwt.sign"))
}

// Generate generate jwt token from claims
func Generate(claims jwt.Claims) *jwt.Token {
	return jwt.NewWithClaims(method, claims)
}

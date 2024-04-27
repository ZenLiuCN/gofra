package http

import (
	"github.com/ZenLiuCN/gofra/conf"
	"github.com/golang-jwt/jwt/v5"
)

var (
	method jwt.SigningMethod
)

func JwtInitalize() {
	ConfigJwt(conf.GetConfig())
}
func ConfigJwt(conf conf.Config) {
	method = jwt.GetSigningMethod(conf.GetString("jwt.sign"))
}
func Generate(claims jwt.Claims) *jwt.Token {
	return jwt.NewWithClaims(method, claims)
}

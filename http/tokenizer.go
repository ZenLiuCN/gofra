package http

import (
	"github.com/ZenLiuCN/goinfra/conf"
	"github.com/golang-jwt/jwt/v5"
)

var (
	method jwt.SigningMethod
)

func init() {
	ConfigJwt(conf.GetConfigurer())
}
func ConfigJwt(conf conf.Config) {
	method = jwt.GetSigningMethod(conf.GetString("jwt.sign"))
}
func Generate(claims jwt.Claims) *jwt.Token {
	return jwt.NewWithClaims(method, claims)
}

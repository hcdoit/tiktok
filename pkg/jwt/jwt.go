package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/hcdoit/tiktok/pkg/consts"
	"github.com/hcdoit/tiktok/pkg/errno"
)

// private claims, share information between parties that agree on using them
// CustomClaims Structured version of Claims Section, as referenced at https://tools.ietf.org/html/rfc7519#section-4.1 See examples for how to use this with your own claim types
type CustomClaims struct {
	Id          int64
	AuthorityId int64
	jwt.StandardClaims
}

// CreateToken creates a new token
func CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// zap.S().Debugf(token.SigningString())
	return token.SignedString([]byte(consts.SecretKey))

}

// ParseToken parses the token.
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(consts.SecretKey), nil
	})
	if err != nil {
		return nil, errno.AuthInvalidJwt
	}
	// verify the token claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errno.AuthInvalidJwt
}

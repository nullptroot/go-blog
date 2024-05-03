package userser

import (
	redisser "go-blog/service/redis_ser"
	"go-blog/utils/jwts"
	"time"
)

func (UserService) Logout(claims *jwts.CustomClaims, token string) error {
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	return redisser.Logout(token, diff)
}

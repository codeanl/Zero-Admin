package jwt

import "github.com/golang-jwt/jwt/v4"

func GetJwtToken(secretKey string, iat, seconds, id int64, username string, nickname string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["id"] = id
	claims["username"] = username
	claims["nickname"] = nickname
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
func GetMemberJwtToken(secretKey string, iat, seconds, MemberId int64, username string, nickname string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["memberId"] = MemberId
	claims["username"] = username
	claims["nickname"] = nickname
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

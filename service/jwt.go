package service

import (
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/ed25519"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Claims struct {
	Dom string `json:"dom"`
	jwt.RegisteredClaims
}

type JWTPayloadData struct {
	Iss string
	Sub string
	Aud string
	Id  string
}

var PrivateKey = os.Getenv("POCKAMENT_PRIVATEKEY_PATH")
var PublicKey = os.Getenv("POCKAMENT_PUBLICKEY_PATH")
var DomainName = os.Getenv("POCKAMENT_DOMAIN")

func GenJwtToken(option JWTPayloadData) string {

	// 秘密鍵ファイルを読み込む
	file, err := ioutil.ReadFile(PrivateKey)
	if err != nil {
		return ""
	}

	// claimを書き込む
	claims := Claims{
		DomainName,
		jwt.RegisteredClaims{
			Issuer:    option.Iss,
			Subject:   option.Sub,
			Audience:  []string{option.Aud},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        option.Id,
		}}

	// Headerとclaimを書き込み
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	// 鍵はbase64でエンコードされているのでデコード
	dec, _ := base64.StdEncoding.DecodeString(string(file))
	// 型を明示しておく
	var DecodedKey ed25519.PrivateKey
	DecodedKey = dec
	// 署名
	ss, err := token.SignedString(DecodedKey)
	if err != nil {
		return ""
	}
	return ss
}

// CheckJWTToken TokenのClaim部分を返す
func CheckJWTToken(token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		pk, err := ioutil.ReadFile(PublicKey)
		dec, _ := base64.StdEncoding.DecodeString(string(pk))
		var Pkey ed25519.PublicKey = dec
		if err != nil {
			return nil, errors.New("failed to read key")
		}

		if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
			err := errors.New("unexpected signing method")
			return nil, err
		}
		return Pkey, nil
	})
	if err != nil {
		return "", err
	}
	if !parsedToken.Valid {
		return "", errors.New("token is invalid")
	}

	s := strings.Split(parsedToken.Raw, ".")
	decodeString, err := base64.RawURLEncoding.DecodeString(s[1])
	if err != nil {
		return "", err
	}
	return string(decodeString), nil

}

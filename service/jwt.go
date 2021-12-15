package service

import (
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/ed25519"
	"io/ioutil"
	"os"
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

var PrivateKey = os.Getenv("OSHAVERY_PRIVATEKEY_PATH")
var DomainName = os.Getenv("OSHAVERY_DOMAIN")

func GenJwtToken(option JWTPayloadData) {

	// 秘密鍵ファイルを読み込む
	file, err := ioutil.ReadFile(PrivateKey)
	if err != nil {
		return
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
		return
	}
	fmt.Printf("%v\n", ss)
}

package service

import (
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/ed25519"
	"io/ioutil"
	"time"
)

type Claims struct {
	Dom string `json:"dom"`
	jwt.RegisteredClaims
}

func GenJwtToken() {
	// 秘密鍵ファイルを読み込む
	file, err := ioutil.ReadFile("p.pem")
	if err != nil {
		return
	}

	// claimを書き込む
	claims := Claims{
		"auth.p0at.nanai10a.net",
		jwt.RegisteredClaims{
			Issuer:    "sign@auth.p0at.nanai10a.net",
			Subject:   "general",
			Audience:  []string{"laminne33569"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        "574c3dc7-94c1-46ee-8b52-b7e6fdcd4e1f",
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

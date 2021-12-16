/*
Pockament-auth
(C) 2021 Tatsuto Yamamoto
*/
package main

import (
	"fmt"
	"github.com/pockament/auth/repository"
	"github.com/pockament/auth/routers"
	"github.com/pockament/auth/service"
	"os"
)

func main() {
	if os.Args[1] != "test" {
		token := service.GenJwtToken(service.JWTPayloadData{
			Iss: "lami",
			Sub: "test",
			Aud: "aa",
			Id:  "33559fc5-5dc8-11ec-b08f-80fa5b7866d5",
		})
		//jwtToken, _ := service.CheckJWTToken(token)
		fmt.Printf("%v\n", token)

	} else {
		repository.RepositoryMain()
		routers.Serve()
	}
}

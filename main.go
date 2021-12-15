/*
Pockament-auth
(C) 2021 Tatsuto Yamamoto
*/
package main

import (
	"github.com/pockament/auth/repository"
	"github.com/pockament/auth/routers"
)

func main() {
	repository.RepositoryMain()
	routers.Serve()
}

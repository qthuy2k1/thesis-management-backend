package handler

import (
	"log"

	oauth2pkg "github.com/qthuy2k1/thesis-management-backend/user-svc/pkg/oauth2"
)

func GoogleLogin() {
	googleConfig := oauth2pkg.SetupConfig()

	log.Println(googleConfig)
}

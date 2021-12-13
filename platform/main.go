package main

import (
	_ "github.com/FuradWho/ChaincodeDeployment/platform/third_party/logger"
	"github.com/FuradWho/ChaincodeDeployment/platform/web/controllers"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.Infoln("dasdsad")

	controllers.StartIris()

}

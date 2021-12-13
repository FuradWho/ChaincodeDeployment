package test

import (
	"fmt"
	"github.com/FuradWho/ChaincodeDeployment/platform/models"
	"testing"

	_ "github.com/FuradWho/ChaincodeDeployment/platform/third_party/logger"
)

func TestClient(t *testing.T) {

	fabricClient := new(models.FabricClient)

	err := fabricClient.Init()
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	fmt.Printf("%+v", fabricClient.NetworkConfig)
}

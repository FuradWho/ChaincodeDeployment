package models

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config/lookup"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

const (
	channelId        = "mychannel"
	channelTx        = "/usr/local/hyper/test2/configtx/channel-artifacts/mychannel.tx"
	connectConfigDir = "/home/fabric/GolandProjects/ChaincodeDeployment/platform/configs/connect-config/sdk-connection-config.yaml"
	chaincodeId      = "mycc_0"
	chaincodePath    = "newchaincode/test"
	ccVersion        = "0"
	Admin            = "Admin"
)

type FabricClient struct {
	ConnectionFile []byte
	NetworkConfig  fab.NetworkConfig
	ChannelId      string
	GoPath         string

	userName string
	userOrg  string

	resmgmtClients map[string]*resmgmt.Client
	sdk            *fabsdk.FabricSDK
	retry          resmgmt.RequestOption
}

func (f *FabricClient) GetNetworkConfig() (fab.NetworkConfig, error) {

	configBackend, _ := f.sdk.Config()
	networkConfig := fab.NetworkConfig{}

	err := lookup.New(configBackend).UnmarshalKey("organizations", &networkConfig.Organizations)
	if err != nil {
		log.Errorf("Failed to unmarsha org :%s \n", err)
		return networkConfig, err
	}

	err = lookup.New(configBackend).UnmarshalKey("orderers", &networkConfig.Orderers)
	if err != nil {
		log.Errorf("Failed to unmarsha org :%s \n", err)
		return networkConfig, err
	}

	err = lookup.New(configBackend).UnmarshalKey("channels", &networkConfig.Channels)
	if err != nil {
		log.Errorf("Failed to unmarsha org :%s \n", err)
		return networkConfig, err
	}

	err = lookup.New(configBackend).UnmarshalKey("peers", &networkConfig.Peers)
	if err != nil {
		log.Errorf("Failed to unmarsha org :%s \n", err)
		return networkConfig, err
	}

	return networkConfig, nil
}

func (f *FabricClient) Init() error {

	connectConfig, _ := ioutil.ReadFile(connectConfigDir)

	f.ConnectionFile = connectConfig
	f.ChannelId = channelId
	f.GoPath = os.Getenv("GOPATH")

	sdk, err := fabsdk.New(config.FromRaw(f.ConnectionFile, "yaml"))
	if err != nil {
		log.Error("Failed to setup main sdk ")
		return err
	}
	f.sdk = sdk

	networkConfig, err := f.GetNetworkConfig()
	if err != nil {
		log.Error("Failed to Get Network Config")
		return err
	}

	resmgmtClients := make(map[string]*resmgmt.Client)
	f.NetworkConfig = networkConfig

	for k, _ := range f.NetworkConfig.Organizations {
		resmgmtClient, err := resmgmt.New(sdk.Context(fabsdk.WithUser(Admin), fabsdk.WithOrg(k)))
		if err != nil {
			log.Errorf("Failed to create channel management client : %s \n", err)
			return err
		}
		resmgmtClients[k] = resmgmtClient
	}

	f.resmgmtClients = resmgmtClients
	f.retry = resmgmt.WithRetry(retry.DefaultResMgmtOpts)

	return nil
}

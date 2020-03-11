package sim

import (
	"github.com/irisnet/irishub-sdk-go/client"
	"github.com/irisnet/irishub-sdk-go/crypto"
	"github.com/irisnet/irishub-sdk-go/types"
)

const (
	NodeURI  = "localhost:26657"
	ChainID  = "irishub-test"
	Online   = true
	Network  = types.Testnet
	Mode     = types.Commit
	Fee      = "600000000000000000iris-atto"
	Gas      = 20000
	mnemonic = "small culture theory rare offer polar seek mule planet fog garlic segment burger guard bar cool milk lion loyal head olympic impulse purpose forget"
)

var (
	priKey string
	addr   string
)

type TestClient struct {
	sender types.AccAddress
	client.SDKClient
}

func NewClient() TestClient {
	keyManager, err := crypto.NewMnemonicKeyManager(mnemonic)
	if err != nil {
		panic(err)
	}

	priKey, err = keyManager.ExportAsPrivateKey()
	if err != nil {
		panic(err)
	}
	addr = types.AccAddress(keyManager.GetPrivKey().PubKey().Address()).String()
	fees, err := types.ParseCoins(Fee)
	if err != nil {
		panic(err)
	}

	client := client.NewSDKClient(types.SDKConfig{
		NodeURI:   NodeURI,
		Network:   Network,
		ChainID:   ChainID,
		Gas:       Gas,
		Fee:       fees,
		KeyDAO:    createTestKeyDAO(),
		Mode:      Mode,
		Online:    Online,
		StoreType: types.Keystore,
		Level:     "debug",
	})
	return TestClient{
		SDKClient: client,
		sender:    types.MustAccAddressFromBech32(addr),
	}
}

func (tc TestClient) Sender() types.AccAddress {
	return tc.sender
}

func createTestKeyDAO() *TestKeyDAO {
	dao := TestKeyDAO{
		store: map[string]types.Store{},
	}
	keystore := types.KeyInfo{
		PrivKey: priKey,
		Address: addr,
	}
	_ = dao.Write("test1", keystore)
	return &dao
}

type TestKeyDAO struct {
	store map[string]types.Store
}

func (dao *TestKeyDAO) Write(name string, store types.Store) error {
	dao.store[name] = store
	return nil
}

func (dao *TestKeyDAO) Read(name, pwd string) (types.Store, error) {
	return dao.store[name], nil
}

func (dao *TestKeyDAO) Delete(name, pwd string) error {
	delete(dao.store, name)
	return nil
}

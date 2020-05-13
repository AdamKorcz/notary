package cryptoservice

import (
	"github.com/theupdateframework/notary/trustmanager"
	"github.com/theupdateframework/notary/tuf/data"
)

var passphraseRetriever = func(string, string, bool, int) (string, bool, error) { return "", false, nil }

type CryptoServiceTester struct {
	role    data.RoleName
	keyAlgo string
	gun     data.GUN
}

func (c CryptoServiceTester) cryptoServiceFactory() *CryptoService {
	return NewCryptoService(trustmanager.NewKeyMemoryStore(passphraseRetriever))
}

func Fuzz(data []byte) int {
	c := CryptoServiceTester{}
	cryptoService := c.cryptoServiceFactory()
	_, _, err := cryptoService.GetPrivateKey(string(data))
	if err != nil {
		return 0
	}
	return 1
}

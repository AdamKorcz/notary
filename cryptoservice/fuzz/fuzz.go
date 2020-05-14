import (
        "github.com/theupdateframework/notary/trustmanager"
        "github.com/theupdateframework/notary/cryptoservice"
        "github.com/theupdateframework/notary/passphrase"
)

func Fuzz(data []byte) int {
        cryptos := cryptoservice.NewCryptoService(trustmanager.NewKeyMemoryStore(passphrase.ConstantRetriever("pass")))
        _, _, err := cryptos.GetPrivateKey(string(data))
        if err != nil {
                return 0
        }
        return 1
}

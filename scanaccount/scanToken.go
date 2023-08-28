package scanaccount

import (
	"fmt"
	"math/big"

	txnBuilder "github.com/coming-chat/go-aptos/transaction_builder"
)

func (c *TokenClient) GetAllToken(account txnBuilder.AccountAddress) {
	resType, _ := c.GetAccountResources(account.ToShortString(), 0) // 0 means the ;atest version
	for _, res := range resType {
		if StartsWith(res.Type, "0x1::coin::CoinStore") {
			fmt.Println(res.Type)
			coin := res.Data["coin"].(map[string]interface{})
			value := coin["value"].(string)
			balance, _ := big.NewInt(0).SetString(value, 10)
			fmt.Println(balance)
		}
	}
}

func StartsWith(s, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if s[i] != prefix[i] {
			return false
		}
	}
	return true
}

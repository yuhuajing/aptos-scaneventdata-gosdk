package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/coming-chat/go-aptos/aptosclient"
	"github.com/coming-chat/go-aptos/aptostypes"
	"github.com/coming-chat/go-aptos/scanaccount"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const rpcUrl = "https://fullnode.devnet.aptoslabs.com"
const MainnetRestUrl = "https://fullnode.mainnet.aptoslabs.com"

//const rpcUrl = "https://fullnode.mainnet.aptoslabs.com"

func main() {
	ctx := context.Background()
	client, err := aptosclient.Dial(ctx, MainnetRestUrl)
	if err != nil {
		printError(err)
	}

	address := "0x7e5f7bdd454478be1ffe9b66b849efd02359a971aa6a848ceb03bbb5729b3b52"

	dba := scanaccount.Buildconnect()
	dba.AutoMigrate(&scanaccount.NFTInfo{}, &scanaccount.AccToken{}, &scanaccount.NFTOwner{})

	//accountinfo, _ := client.GetAccount(address)

	res := scanaccount.AccountAddress{}
	dba.Model(&scanaccount.AccountAddress{}).Where("address = ?", address).Find(&res)
	// for {
	// 	if int(accountinfo.SequenceNumber) > res.Sequence {
	// 		scanaccount.GetAllTokenForAccount(dba, client, address, res.Sequence)
	// 		scanaccount.GetAllToken(dba, client, address)
	// 		dba.Model(&scanaccount.AccountAddress{}).Where("address = ?", address).Update("sequence", int(accountinfo.SequenceNumber))
	// 	} else {
	// 		time.Sleep(5 * time.Minute)
	// 	}
	// }

	scanaccount.GetAllTokenForAccount(dba, client, address, 0)
	scanaccount.GetAllToken(dba, client, address)
}

func printLine(content string) {
	fmt.Printf("================= %s =================\n", content)
}

func printError(err error) {
	var restError *aptostypes.RestError
	if b := errors.As(err, &restError); b {
		fmt.Printf("code: %d, message: %s, aptos_ledger_version: %d\n", restError.Code, restError.Message, restError.AptosLedgerVersion)
	} else {
		fmt.Printf("err: %s\n", err.Error())
	}
}

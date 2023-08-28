package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/coming-chat/go-aptos/aptosclient"
	"github.com/coming-chat/go-aptos/aptostypes"
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

	ledgerInfo, err := client.LedgerInfo()
	if err != nil {
		panic(err)
	}
	content, err := json.Marshal(ledgerInfo)
	if err != nil {
		printError(err)
	}
	fmt.Println(string(content))

	//resaccount := "0x80e9dd042dd7e48b2325f4383ae02f49b542f106dbcdec689fdfe2f3c83c67d3"
	getAccount(client, "0x855b0acf4fe5594272fcc258524718a83cd4e6e457f7af2c32b6e962d3516d3")
	//getAccountResourceByTxNumber(client, "0x31b4c2529a4fcf3cc128e6d28fa1becd93f68dc05067924dbdf60a2c6077fc0f", ledgerInfo.LedgerVersion)
	//getAccResByTxNumberAndResType(client, resaccount, "0x1::account::Account")
	//getAptosBalance(client, "0x28418760527ca583ca1dea9807b21574952f9eac6f2cbef9369f0c87f1940497")
	//getBalanceByCoinType(client, "0xe8ca094fec460329aaccc2a644dc73c5e39f1a2ad6e97f82b6cbdc1a5949b9ea", "0xcc78307c77f1c2c0fdfee17269bfca7876a0b35438c3442417480c0d5c370fbc::AptopadCoin::APD")
	//moduleaccount := "0x8d2d7bcde13b2513617df3f98cdd5d0e4b9f714c6308b9204fe18ad900d92609"
	//getAccountModulesByTxNumber(client, moduleaccount)
	//getAccModByTxNumberAndModName(client, moduleaccount, "mint")
	//getBlockByHeight(client, 79536591, false)
	//getBlockByVersion(client, 218980864, true)
	//getEventsByCreationNumber(client, "0x92d2f7ad00630e4dfffcca01bee12c84edf004720347fb1fd57016d2cc8d3f8", 10, 305200, 2)
	//getEventsByHandle(client, "0x1f92870c5f3db310fcd81e2a46c1e1ddb576049837294e46f43c2b79cf8213c0", "0x1::coin::CoinStore<0x1::aptos_coin::AptosCoin>", "withdraw_events", 1, 5)

	// transactions, err := client.GetTransactions(219032110, 1)
	// if err != nil {
	// 	printError(err)
	// }
	// printLine("get tx list")
	// for _, tx := range transactions {
	// 	fmt.Printf("sender %s, version: %d, type: %s, hash: %s, sequeneNumber: %d, time: %d\n", tx.Sender, tx.Version, tx.Type, tx.Hash, tx.SequenceNumber, tx.Timestamp)
	// }

	// tx, _ := client.GetTransactionByHash("0x7499f8c9d2f9cf978b314e384b44b9bc8fa939bdc0d627891be11b543a64c25b")
	// content, err := json.Marshal(tx)
	// if err != nil {
	// 	printError(err)
	// }
	// fmt.Println(string(content))
	// price, _ := client.EstimateGasPrice()
	// fmt.Println(price)

	// tx, _ := client.GetTransactionByHash("0xbb462783c5f7b252c7a3502f9c9f58c0f6918bfb5f95c9c37cf5eaa66c4b9847")
	// fmt.Println(tx.Sender)
	// fmt.Println(tx.SequenceNumber)

	// events, err := client.GetEventsByCreationNumber("0xf980d8d14637dfaf9b961b6b717dad814a79d667f038d9f534bebf07b59d68b", "3", 56, 1)
	// if err != nil {
	// 	printError(err)
	// }
	// for _, e := range events {
	// 	fmt.Println(e.SequenceNumber)
	// 	fmt.Println(e.Data)
	// }
}
func getEventsByHandle(client *aptosclient.RestClient, account, handle, field string, start, limit uint64) {
	event, _ := client.GetEventsByEventHandle(account, handle, field, start, limit)
	content, err := json.Marshal(event[0])
	if err != nil {
		printError(err)
	}
	fmt.Println(string(content))
}

func getEventsByCreationNumber(client *aptosclient.RestClient, account string, creationNumber int, start, limit uint64) {
	event, _ := client.GetEventsByCreationNumber(account, strconv.Itoa(creationNumber), start, limit)
	content, err := json.Marshal(event[0])
	if err != nil {
		printError(err)
	}
	fmt.Println(string(content))
}
func getBlockByVersion(client *aptosclient.RestClient, height int, withTx bool) {
	gotBlock, _ := client.GetBlockByVersion(strconv.Itoa(height), withTx)
	content, err := json.Marshal(gotBlock)
	if err != nil {
		printError(err)
	}
	fmt.Println(string(content))
}
func getBlockByHeight(client *aptosclient.RestClient, height int, withTx bool) {
	gotBlock, _ := client.GetBlockByHeight(strconv.Itoa(height), withTx)
	content, err := json.Marshal(gotBlock)
	if err != nil {
		printError(err)
	}
	fmt.Println(string(content))
}
func getBalanceByCoinType(client *aptosclient.RestClient, account string, cointype string) {
	coininfo, _ := client.GetCoinInfo(cointype)
	fmt.Println(coininfo)
	value, _ := client.BalanceOf(account, cointype)
	fmt.Println(value.Int64())
}
func getAptosBalance(client *aptosclient.RestClient, account string) {
	value, _ := client.AptosBalanceOf(account)
	fmt.Println(value.Int64())
}
func getAccModByTxNumberAndModName(client *aptosclient.RestClient, account string, modname string) {
	//firstTx, _ := client.GetAccountTransactions(account, 1, 1)
	modres, _ := client.GetAccountModuleByModname(account, modname, 0)
	fmt.Println(modres.Abi.Name)
}
func getAccountModulesByTxNumber(client *aptosclient.RestClient, account string) {
	firstTx, _ := client.GetAccountTransactions(account, 1, 1)
	modres, _ := client.GetAccountModules(account, firstTx[0].Version)
	fmt.Println(modres[0].Abi.Name)
}
func getAccResByTxNumberAndResType(client *aptosclient.RestClient, account string, resType string) {
	firstTx, _ := client.GetAccountTransactions(account, 1, 1)
	accountResource, _ := client.GetAccountResourceByResType(account, resType, firstTx[0].Version)
	fmt.Println(accountResource)
}
func getAccountResourceByTxNumber(client *aptosclient.RestClient, account string, txNumber uint64) {
	//firstTx, _ := client.GetAccountTransactions(account, 1, 1)
	//accountResource, _ := client.GetAccountResources(account, firstTx[0].Version)
	accountResource, _ := client.GetAccountResources(account, txNumber)
	fmt.Println(accountResource)
	//fmt.Println(accountResource[0].Data["authentication_key"])
}
func getAccount(client *aptosclient.RestClient, account string) {
	accountinfo, _ := client.GetAccount(account)
	by, _ := accountinfo.MarshalJSON()
	fmt.Println(string(by))
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

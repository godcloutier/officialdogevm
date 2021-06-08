/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.


The tx creation was taken from https://www.thepolyglotdeveloper.com/2018/03/create-sign-bitcoin-transactions-golang/


*/
package cmd

import (
	"fmt"
	"bytes"
	"encoding/hex"
	
	"github.com/spf13/cobra"
//	"github.com/btcsuite/btcd/chaincfg"
	//"github.com/rosetta-dogecoin/rosetta-dogecoin/bitcoin"
	"github.com/rosetta-dogecoin/rosetta-dogecoin/configuration"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"	
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	//"github.com/rosetta-dogecoin/rosetta-dogecoin/dogecoin/chaincfg"
	"log"
)

// txtestCmd represents the txtest command
var txtestCmd = &cobra.Command{
	Use:   "txtest",
	Short: "creates a tx on the dogechain",
	Long: `basic transaction creation test.`,
	Run: func(cmd *cobra.Command, args []string) {


	fmt.Println("usage: dogity txtest <balance,build>")
		
	},
}

// txtestCmd represents the config command
var txtestBalanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "shows balance",
	Long: `connection test to dogecoind, shows account balances.`,
	Run: func(cmd *cobra.Command, args []string) {

	fmt.Println("[txtest balance] called")
	balanceTest();
		
	},
}



// txtestCmd represents the config command
var txtestBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "build transaction",
	Long: `basic transaction creation test.`,
	Run: func(cmd *cobra.Command, args []string) {

	fmt.Println("[txtest build] called")
	balanceTest();
		
	},
}

func buildTest() {

	// create new client instance
	client, err := rpcclient.New(&rpcclient.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   true,
		Host:         "127.0.0.1:22555",
		User:         "smartdoge",
		Pass:         "verySmartMuchWow!",
	}, nil)
	if err != nil {
		log.Fatalf("error creating new doge client: %v", err)
	}


	fmt.Println(chaincfg)


	// list accounts
	accounts, err := client.ListAccounts()
	if err != nil {
		log.Fatalf("error listing accounts: %v", err)
	}



	// iterate over accounts (map[string]btcutil.Amount) and write to stdout
	for label, amount := range accounts {
		log.Printf("%s: %s", label, amount)
	}

}


type MyTransaction struct {
	TxId               string `json:"txid"`
	SourceAddress      string `json:"source_address"`
	DestinationAddress string `json:"destination_address"`
	Amount             int64  `json:"amount"`
	UnsignedTx         string `json:"unsignedtx"`
	SignedTx           string `json:"signedtx"`
}


func CreateTx(secret string, destination string, amount int64, txHash string) (MyTransaction, error) {
	var transaction MyTransaction
	wif, err := btcutil.DecodeWIF(secret)
	if err != nil {
		return MyTransaction{}, err
	}
	fmt.Println(&chaincfg)
	addresspubkey, _ := btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeUncompressed(), &chaincfg.MainNetParams)
	sourceTx := wire.NewMsgTx(wire.TxVersion)
	sourceUtxoHash, _ := chainhash.NewHashFromStr(txHash)
	sourceUtxo := wire.NewOutPoint(sourceUtxoHash, 0)
	sourceTxIn := wire.NewTxIn(sourceUtxo, nil, nil)
	destinationAddress, err := btcutil.DecodeAddress(destination, &chaincfg.MainNetParams)
	sourceAddress, err := btcutil.DecodeAddress(addresspubkey.EncodeAddress(), &chaincfg.MainNetParams)
	if err != nil {
		return MyTransaction{}, err
	}
	destinationPkScript, _ := txscript.PayToAddrScript(destinationAddress)
	sourcePkScript, _ := txscript.PayToAddrScript(sourceAddress)
	sourceTxOut := wire.NewTxOut(amount, sourcePkScript)
	sourceTx.AddTxIn(sourceTxIn)
	sourceTx.AddTxOut(sourceTxOut)
	sourceTxHash := sourceTx.TxHash()
	redeemTx := wire.NewMsgTx(wire.TxVersion)
	prevOut := wire.NewOutPoint(&sourceTxHash, 0)
	redeemTxIn := wire.NewTxIn(prevOut, nil, nil)
	redeemTx.AddTxIn(redeemTxIn)
	redeemTxOut := wire.NewTxOut(amount, destinationPkScript)
	redeemTx.AddTxOut(redeemTxOut)
	sigScript, err := txscript.SignatureScript(redeemTx, 0, sourceTx.TxOut[0].PkScript, txscript.SigHashAll, wif.PrivKey, false)
	if err != nil {
		return MyTransaction{}, err
	}
	redeemTx.TxIn[0].SignatureScript = sigScript
	flags := txscript.StandardVerifyFlags
	vm, err := txscript.NewEngine(sourceTx.TxOut[0].PkScript, redeemTx, 0, flags, nil, nil, amount)
	if err != nil {
		return MyTransaction{}, err
	}
	if err := vm.Execute(); err != nil {
		return MyTransaction{}, err
	}
	var unsignedTx bytes.Buffer
	var signedTx bytes.Buffer
	sourceTx.Serialize(&unsignedTx)
	redeemTx.Serialize(&signedTx)
	transaction.TxId = sourceTxHash.String()
	transaction.UnsignedTx = hex.EncodeToString(unsignedTx.Bytes())
	transaction.Amount = amount
	transaction.SignedTx = hex.EncodeToString(signedTx.Bytes())
	transaction.SourceAddress = sourceAddress.EncodeAddress()
	transaction.DestinationAddress = destinationAddress.EncodeAddress()
	return transaction, nil
}



func balanceTest() {

	// create new client instance
	client, err := rpcclient.New(&rpcclient.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   true,
		Host:         "127.0.0.1:22555",
		User:         "smartdoge",
		Pass:         "verySmartMuchWow!",
	}, nil)
	if err != nil {
		log.Fatalf("error creating new doge client: %v", err)
	}

	// list accounts
	accounts, err := client.ListAccounts()
	if err != nil {
		log.Fatalf("error listing accounts: %v", err)
	}
	// iterate over accounts (map[string]btcutil.Amount) and write to stdout
	for label, amount := range accounts {
		log.Printf("%s: %s", label, amount)
	}

/*
	// prepare a sendMany transaction
	receiver1, err := btcutil.DecodeAddress("1someAddressThatIsActuallyReal", &chaincfg.MainNetParams)
	if err != nil {
		log.Fatalf("address receiver1 seems to be invalid: %v", err)
	}
	receiver2, err := btcutil.DecodeAddress("1anotherAddressThatsPrettyReal", &chaincfg.MainNetParams)
	if err != nil {
		log.Fatalf("address receiver2 seems to be invalid: %v", err)
	}
	receivers := map[btcutil.Address]btcutil.Amount{
		receiver1: 42,  // 42 satoshi
		receiver2: 100, // 100 satoshi
	}
*/
	
	// create and send the sendMany tx
	//txSha, err := client.SendMany("some-account-label-from-which-to-send", receivers)
	//if err != nil {
	//	log.Fatalf("error sendMany: %v", err)
	//}
	//log.Printf("sendMany completed! tx sha is: %s", txSha.String())

}


func init() {
	rootCmd.AddCommand(txtestCmd)
	txtestCmd.AddCommand(txtestBalanceCmd)
	txtestCmd.AddCommand(txtestBuildCmd)
	// Here you will define your flags and chaincfg settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// txtestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// txtestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

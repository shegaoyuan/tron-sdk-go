package main

import (
	"encoding/hex"
	"fmt"
	"github.com/bytejedi/tron-sdk-go/abi"
	"github.com/bytejedi/tron-sdk-go/utils"
	"log"

	"github.com/bytejedi/tron-sdk-go/keystore"

	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
	"github.com/tyler-smith/go-bip39/wordlists"
)

func main() {
	keystorePassword := "password"

	// Set wordlist
	bip39.SetWordList(wordlists.ChineseSimplified)
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		log.Fatal(err)
	}

	mnemonic, _ := bip39.NewMnemonic(entropy)
	seed := bip39.NewSeed(mnemonic, keystorePassword)

	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/195'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	privKeyHex, err := wallet.PrivateKeyHex(account)
	if err != nil {
		log.Fatal(err)
	}

	pubKeyHex, err := wallet.PublicKeyHex(account)
	if err != nil {
		log.Fatal(err)
	}

	privateKeyECDSA, err := wallet.PrivateKey(account)
	if err != nil {
		log.Fatal(err)
	}
	publicKeyECDSA, err := wallet.PublicKey(account)
	if err != nil {
		log.Fatal(err)
	}

	keystoreKey := keystore.NewKeyFromECDSA(privateKeyECDSA)
	keyjson, err := keystore.EncryptKey(keystoreKey, keystorePassword, keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		log.Fatal(err)
	}

	tronAddress := keystore.PubkeyToAddress(*publicKeyECDSA)

	fmt.Println("助记词:", mnemonic)
	fmt.Println("base58地址:", tronAddress.String())
	fmt.Println("hex地址:", hex.EncodeToString(tronAddress))
	fmt.Println("私钥:", privKeyHex)
	fmt.Println("公钥:", pubKeyHex)
	fmt.Println("keystore:", string(keyjson))

	paramStr := "[{\"address\":\"TRu2DruRJDjVsqno7CwXMzJb7vQTpVaKmL\"},{\"address\":\"TRu2DruRJDjVsqno7CwXMzJb7vQTpVaKmL\"},{\"uint256\":\"10000\"},{\"uint256\":\"0\"}]"
	param, err := abi.LoadFromJSON(paramStr)
	if err != nil {
		log.Fatal(err)
	}

	paddedParamBytes, err := abi.GetPaddedParam(param)
	if err != nil {
		log.Fatal(err)
	}

	paddedParamHex := utils.Bytes2Hex(paddedParamBytes)
	fmt.Println("triggersmartcontract接口parameter入参:", paddedParamHex)
}

//b3f1c93d000000000000000000000000aeb759c19724572e31d4bc8f7d9d5f3161b056ca000000000000000000000000aeb759c19724572e31d4bc8f7d9d5f3161b056ca00000000000000000000000000000000000000000000000000000000000027100000000000000000000000000000000000000000000000000000000000000000
//000000000000000000000000aeb759c19724572e31d4bc8f7d9d5f3161b056ca000000000000000000000000aeb759c19724572e31d4bc8f7d9d5f3161b056ca00000000000000000000000000000000000000000000000000000000000027100000000000000000000000000000000000000000000000000000000000000000

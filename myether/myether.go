package myether

import (
	"context"
	"fmt"
	Authcid "item/authcid"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var URL = "https://sepolia.infura.io/v3/ddf668d7e38f48c5b62137b857d362a2"               // use the sepolia testnet node
var contractAddress = common.HexToAddress("0x5ba98214fcbb0a21f27d3cda67bc88d61f6585dd") //the address of deployed smart contract

// return the store transaction hash
func Storecid(pk string, cid string) string {
	client, err := ethclient.Dial(URL)
	if err != nil {
		fmt.Printf("ether connect wrong:%v\n", err)
	} else {
		fmt.Println("ether node connected successfully")
	}
	//connect to the url------------------------------
	cnt := context.Background()
	chainid, err := client.ChainID(cnt)
	if err != nil {
		fmt.Println("get chainid wrong:", err)
	}
	fmt.Println("chainid is: ", chainid)
	//get chainid------------------------------
	sol, err := Authcid.NewAuthcid(contractAddress, client)
	if err != nil {
		fmt.Println("connect contract wrong:", err)
	}
	sk, _ := crypto.HexToECDSA(pk)
	auth, err := bind.NewKeyedTransactorWithChainID(sk, chainid)
	if err != nil {
		fmt.Println("auth wrong:", err)
	} else {
		fmt.Println("auth success")
	}
	//get the auth of transaction------------------------------
	tx, err := sol.Store(auth, cid)
	if err != nil {
		fmt.Println("excuate store wrong:", err)
	}
	fmt.Println("the store tx is :", tx.Hash().Hex())
	//excuate the store func in solidity------------------------------
	return tx.Hash().Hex()
}

// return the author of cid
func Querycid(cid string) string {
	client, err := ethclient.Dial(URL)
	if err != nil {
		fmt.Printf("ether connect wrong:%v\n", err)
	} else {
		fmt.Println("ether node connected successfully")
	}
	sol, err := Authcid.NewAuthcid(contractAddress, client)
	if err != nil {
		fmt.Println("connect contract wrong:", err)
	}
	author, err := sol.Queryauth(nil, cid)
	if err != nil {
		fmt.Println("query wrong:", err)
	}
	addr := author.Hex() //convert common.address to string
	fmt.Println("the author of cid you queried is:", addr)
	return addr
}

package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"strings"
)

type Block struct {
	previousHash string
	transactions []string
	blockHash    string
}

func newBlock(previousHash string, transactions []string) *Block {

	b := new(Block)
	b.previousHash = previousHash
	b.transactions = transactions

	h := hmac.New(sha256.New, []byte(b.previousHash))
	io.WriteString(h, strings.Join(transactions[:], ""))
	b.blockHash = fmt.Sprintf("%x", h.Sum(nil))

	return b
}

func main() {

	var blockchain []Block

	genesisTransactions := []string{"1000 buks at my wallet"}
	genesisBlock := newBlock("", genesisTransactions)
	blockchain = append(blockchain, *genesisBlock)

	secondTransactions := []string{"2000 buks at my wallet, 1000 buks at John's wallet"}
	secondBlock := newBlock(blockchain[0].blockHash, secondTransactions)
	blockchain = append(blockchain, *secondBlock)

	fmt.Println(blockchain[0])
	fmt.Println(blockchain[1])

}

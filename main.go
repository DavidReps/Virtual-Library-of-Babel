package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

//the chain itself
type BlockChain struct{
	blocks []*Block
}

//blocks themselves
type Block struct{
	Hash []byte
	Data []byte
	PrevHash []byte
}

//will create a hash simple hash function to be replaced later
func (b *Block) DeriveHash(){
	//hash is based on current data and previous hash
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum224(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte)*Block{
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func Genisis()*Block{
	return CreateBlock("Genisis", []byte{})
}

func (chain *BlockChain) AddBlock(data string){
	previous := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, previous.Hash)
	chain.blocks = append(chain.blocks, new)

}

func InitBlockChain() *BlockChain{
	return &BlockChain{[]*Block{Genisis()}}
}

func main(){
	chain := InitBlockChain()

	chain.AddBlock("test for out blockchain")
	chain.AddBlock("second")

	for _, block := range chain.blocks{
		fmt.Printf("hash %x\n", block.PrevHash)
		fmt.Printf("data %s\n", block.Data)
	}
}

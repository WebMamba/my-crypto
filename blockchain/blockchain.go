package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Data []byte
	PrevHash []byte
	Hash []byte
}

func (block *Block) GenerateHash() {
	blockContent := bytes.Join([][]byte{block.PrevHash, block.Data}, []byte{})

	hash := sha256.Sum256(blockContent)

	block.Hash = hash[:]
}

func (chain *BlockChain) AddBlock(data string) {
	block := chain.CreateBlock([]byte(data))

	chain.Blocks = append(chain.Blocks, block)
}

func (chain *BlockChain) CreateBlock(data []byte) (block *Block){
	lastBlock := chain.GetLastBlock()

	block = &Block{Data: data, PrevHash: lastBlock.Hash}
	block.GenerateHash();

	return block
}

func (chain *BlockChain) GetLastBlock() (block *Block) {
	return chain.Blocks[len(chain.Blocks) - 1]
}

func Genesis() (block *Block) {
	firstBlock := Block{Data: []byte("First"), PrevHash: []byte{}}
	firstBlock.GenerateHash()

	return &firstBlock
}

func InitBlockChain() (blockChain *BlockChain) {
	return &BlockChain{[]*Block{Genesis()}}
}
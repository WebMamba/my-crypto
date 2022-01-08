package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"math/big"
	"time"
)

const difficulty = 12

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Data []byte
	PrevHash []byte
	Hash []byte
	Time []byte
	Nonce int64
}

func GenerateHash(block Block, nonce []byte) (hash [32]byte) {
	blockContent := bytes.Join([][]byte{block.PrevHash, block.Data, block.Time, nonce}, []byte{})

	hash = sha256.Sum256(blockContent)

	return hash
}

func (chain *BlockChain) AddBlock(data string) {
	block := chain.CreateBlock([]byte(data))

	chain.Blocks = append(chain.Blocks, block)
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

func (chain *BlockChain) CreateBlock(data []byte) (block *Block) {
	var intHash big.Int
	var currentHash [32]byte

	lastBlock := chain.GetLastBlock()

	block = &Block{Data: data, PrevHash: lastBlock.Hash, Time: []byte(time.Now().String())}

	target := big.NewInt(1)
	target.Lsh(target, uint(256 - difficulty))

	nonce := int64(0)

	for {
		currentHash = GenerateHash(*block, ToHex(nonce))
		intHash.SetBytes(currentHash[:])

		if intHash.Cmp(target) == -1 {
			break
		}

		nonce++
	}

	block.Hash = currentHash[:]
	block.Nonce = nonce

	return block
}

func (block *Block) Validate() bool {
	var intHash big.Int

	nonce := block.Nonce
	hash := GenerateHash(*block, ToHex(nonce))

	target := big.NewInt(1)
	target.Lsh(target, uint(256 - difficulty))

	intHash.SetBytes(hash[:])

	return intHash.Cmp(target) == -1
}

func (chain *BlockChain) GetLastBlock() (block *Block) {
	return chain.Blocks[len(chain.Blocks) - 1]
}

func Genesis() (block *Block) {
	firstBlock := Block{Data: []byte("First"), PrevHash: []byte{}}
	firstBlock.Hash = []byte("hashhhh")

	return &firstBlock
}

func InitBlockChain() (blockChain *BlockChain) {
	return &BlockChain{[]*Block{Genesis()}}
}
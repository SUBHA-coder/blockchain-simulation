package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Block represents a single block in the blockchain
type Block struct {
	Index        int
	Timestamp    string
	Transactions string
	PreviousHash string
	Hash         string
}

// NewBlock creates and returns a new Block
func NewBlock(index int, transactions string, previousHash string) Block {
	block := Block{
		Index:        index,
		Timestamp:    time.Now().String(),
		Transactions: transactions,
		PreviousHash: previousHash,
	}
	block.Hash = block.calculateHash()
	return block
}

// calculateHash calculates the hash of the block
func (b *Block) calculateHash() string {
	record := string(b.Index) + b.Timestamp + b.Transactions + b.PreviousHash
	hash := sha256.New()
	hash.Write([]byte(record))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}

const difficulty = 4 // Number of leading zeros in the hash

// ProofOfWork finds a valid hash for the block by trying different values
func (b *Block) ProofOfWork() {
	for {
		if isValidHash(b.Hash) {
			break
		}
		b.Index++
		b.Hash = b.calculateHash()
	}
}

// isValidHash checks if the block's hash meets the difficulty criteria
func isValidHash(hash string) bool {
	return hash[:difficulty] == "0000"
}

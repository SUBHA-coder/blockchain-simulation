package blockchain

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Blockchain represents the entire blockchain
type Blockchain struct {
	Blocks []Block
}

// NewBlockchain initializes the blockchain with a genesis block
func NewBlockchain() *Blockchain {
	genesisBlock := NewBlock(0, "Genesis Block", "")
	return &Blockchain{Blocks: []Block{genesisBlock}}
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(transactions string) {
	index := len(bc.Blocks)
	previousBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(index, transactions, previousBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// SaveToCSV saves the blockchain data to a CSV file
func (bc *Blockchain) SaveToCSV(filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header only if the file is empty (first run)
	if isEmpty(file) {
		writer.Write([]string{"Index", "Timestamp", "Transactions", "Previous Hash", "Hash"})
	}

	// Write blocks
	for _, block := range bc.Blocks {
		record := []string{
			fmt.Sprintf("%d", block.Index),
			block.Timestamp,
			block.Transactions,
			block.PreviousHash,
			block.Hash,
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("could not write record to csv: %v", err)
		}
	}
	return nil
}

// Helper function to check if the CSV file is empty
func isEmpty(file *os.File) bool {
	stats, _ := file.Stat()
	return stats.Size() == 0
}

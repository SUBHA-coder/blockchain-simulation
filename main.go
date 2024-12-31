package main

import (
	"blockchain-simulation/blockchain"
	"fmt"
	"log"
)

func main() {
	// Initialize the blockchain
	bc := blockchain.NewBlockchain()

	// Add some blocks with transactions
	bc.AddBlock("First block after genesis")
	bc.AddBlock("Second block after genesis")

	// Save blockchain data to CSV after adding blocks
	err := bc.SaveToCSV("blockchain_data.csv")
	if err != nil {
		log.Fatalf("Error saving blockchain to CSV: %v", err)
	}

	// Display all blocks in the blockchain
	for _, block := range bc.Blocks {
		fmt.Printf("Block #%d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Transactions: %s\n", block.Transactions)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println()
	}

	// Add a new block with different transactions and save again
	bc.AddBlock("Third block after genesis")
	err = bc.SaveToCSV("blockchain_data.csv")
	if err != nil {
		log.Fatalf("Error saving blockchain to CSV: %v", err)
	}
}

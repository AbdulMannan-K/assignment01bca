package main

import (
	"fmt"

	BlockChain "github.com/AbdulMannan-K/assignment01bca"
)

func main() {
	// Create a new blockchain and add some blocks.
	blockchain := &BlockChain.Blockchain{}
	blockchain.NewBlock("Mannan to Saad", 64534, "Genesis")
	blockchain.NewBlock("Saad to Awais", 324425, blockchain.Blocks[len(blockchain.Blocks)-1].CurrentHash)
	blockchain.NewBlock("Awais to Tunn", 45634, blockchain.Blocks[len(blockchain.Blocks)-1].CurrentHash)

	// Display all blocks.
	blockchain.DisplayBlocks()

	// Change the transaction of a block.
	blockchain.ChangeBlock(1, "Charlie to Dave")

	// Display all blocks again.
	blockchain.DisplayBlocks()

	// Verify the integrity of the blockchain.
	if blockchain.VerifyChain() {
		fmt.Println("Blockchain is valid.")
	} else {
		fmt.Println("Blockchain is invalid.")
	}
}

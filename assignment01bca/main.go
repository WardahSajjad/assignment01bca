package main

import (
    "crypto/sha256"
    "fmt"
)

// Block represents a block in the blockchain.
type Block struct {
    Transaction   string
    Nonce         int
    PreviousHash  string
    CurrentHash   string
}

// Blockchain is a simple blockchain with a slice of blocks.
type Blockchain struct {
    Blocks []Block
}

// NewBlock creates a new block and adds it to the blockchain.
func (bc *Blockchain) NewBlock(transaction string, nonce int) {
    previousBlock := bc.GetLatestBlock()
    previousHash := previousBlock.CurrentHash
    currentHash := CreateHash(transaction, nonce, previousHash)
    block := Block{
        Transaction:  transaction,
        Nonce:        nonce,
        PreviousHash: previousHash,
        CurrentHash:  currentHash,
    }
    bc.Blocks = append(bc.Blocks, block)
}

// DisplayBlocks prints all the blocks in the blockchain.
func (bc *Blockchain) DisplayBlocks() {
    for i, block := range bc.Blocks {
        fmt.Printf("Block %d:\n", i)
        fmt.Printf("Transaction: %s\n", block.Transaction)
        fmt.Printf("Nonce: %d\n", block.Nonce)
        fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
        fmt.Printf("Current Hash: %s\n\n", block.CurrentHash)
    }
}

// GetLatestBlock returns the latest block in the blockchain.
func (bc *Blockchain) GetLatestBlock() Block {
    if len(bc.Blocks) == 0 {
        // Genesis block
        return Block{}
    }
    return bc.Blocks[len(bc.Blocks)-1]
}

// ChangeBlock changes the transaction of a given block.
func (bc *Blockchain) ChangeBlock(index int, newTransaction string) {
    if index >= 0 && index < len(bc.Blocks) {
        bc.Blocks[index].Transaction = newTransaction
        bc.Blocks[index].CurrentHash = CreateHash(newTransaction, bc.Blocks[index].Nonce, bc.Blocks[index].PreviousHash)
    }
}

// VerifyChain verifies the integrity of the blockchain.
func (bc *Blockchain) VerifyChain() bool {
    for i := 1; i < len(bc.Blocks); i++ {
        if bc.Blocks[i].PreviousHash != bc.Blocks[i-1].CurrentHash {
            return false
        }
    }
    return true
}

// CreateHash calculates the hash of a block.
func CreateHash(transaction string, nonce int, previousHash string) string {
    data := fmt.Sprintf("%s%d%s", transaction, nonce, previousHash)
    hash := sha256.Sum256([]byte(data))
    return fmt.Sprintf("%x", hash)
}

func main() {
    // Initialize the blockchain variable
    blockchain := &Blockchain{
        Blocks: []Block{}, // Initialize the Blocks slice
    }

    // Add some blocks to the blockchain
    blockchain.NewBlock("Alice to Bob", 123)
    blockchain.NewBlock("Bob to Carol", 456)
    blockchain.NewBlock("Carol to Dave", 789)

    // Display all blocks
    blockchain.DisplayBlocks()

    // Change the transaction of the second block
    blockchain.ChangeBlock(1, "New transaction")

    // Verify the blockchain
    if blockchain.VerifyChain() {
        fmt.Println("Blockchain is valid.")
    } else {
        fmt.Println("Blockchain is not valid.")
    }
}

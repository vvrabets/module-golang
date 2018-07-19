package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	blockData int
	link      *block
	prevHash  [32]byte
}

var head *block

func New() *block {
	return &block{}
}

func addBlock(d int) {
	if head == nil {
		head = New()
		head.blockData = d
		fmt.Println(head.blockData)
		head.prevHash = sha256.Sum256([]byte(""))
		fmt.Println(head.prevHash)
		return
	}
	currentBlock := head
	for currentBlock.link != nil {
		currentBlock = currentBlock.link
	}
	newBlock := New()
	currentBlock.link = newBlock
	newBlock.blockData = d
	fmt.Println(newBlock.blockData)
}

func main() {
	addBlock(20)
	addBlock(40)
	//fmt.Println(head.link.blockData)
	fmt.Println("vim-go")
}

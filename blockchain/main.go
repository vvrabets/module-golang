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
		fmt.Println("head.prewHash:", head.prevHash)
		return
	}
	currentBlock := head
	for currentBlock.link != nil {
		currentBlock = currentBlock.link
	}
	newBlock := New()
	currentBlock.link = newBlock
	newBlock.blockData = d
	fmt.Println("*currentBlock:", currentBlock.prevHash)
	fmt.Println("newBlock.blockData:", newBlock.blockData)
	arr := make([]byte, len(currentBlock.prevHash))
	for i := 0; i < len(currentBlock.prevHash); i++ {
		arr[i] = currentBlock.prevHash[i]
	}
	newBlock.prevHash = sha256.Sum256([]byte(arr))
	fmt.Println("newBlock.prevHash:", newBlock.prevHash)
}

func main() {
	addBlock(20)
	addBlock(40)
	//fmt.Println(head.link.blockData)
	fmt.Println("vim-go")
}

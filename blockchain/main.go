package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
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
		head.prevHash = sha256.Sum256([]byte("zdarova"))
		fmt.Println("head.prevHash:", head.prevHash)
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

func verifyChain() {
	if head == nil {
		fmt.Println("block is empty!Try again after adding some blocks  ")
		return
	}
	curr := head.link
	prev := head
	var count int = 1
	for curr != nil {
		count++
		fmt.Printf("count: %v , Data: %v , Hash: %v \n", count, curr.blockData, curr.prevHash)
		//fmt.Println("prevHash:", cur.prevHash)
		arr := make([]byte, len(prev.prevHash))
		for i := 0; i < len(prev.prevHash); i++ {
			arr[i] = prev.prevHash[i]
		}
		b := sha256.Sum256([]byte(arr))
		if b == curr.prevHash {
			fmt.Println("Verified\n")
		} else {
			fmt.Println("Verification failed\n")
		}
		prev = curr
		curr = curr.link
	}
}

func alterNthBlock(n int, newData int) {
	curr := head
	if curr == nil {
		fmt.Println("Nth block does not exist! \n")
		return
	}
	count := 0
	for count != n {
		if curr.link == nil && count != n {
			fmt.Println("Nth block does not exist\n")
		} else if count == n {
			break
		}
		curr = curr.link
		count++
	}
	fmt.Println("Before: ")
	printBlock(curr)
	curr.blockData = newData
	fmt.Println(" \nAfter")
	printBlock(curr)
}

func hackChain() {
	curr := head
	var prev *block
	if curr == nil {
		fmt.Println("BlockChain is empty ")
		return
	}
	for {
		prev = curr
		curr = curr.link
		if curr == nil {
			return
		}
		arr := make([]byte, len(prev.prevHash))
		for i := 0; i < len(prev.prevHash); i++ {
			arr[i] = prev.prevHash[i]
		}
		b := sha256.Sum256([]byte(arr))
		if b == curr.prevHash {
			fmt.Printf("%v and  %v ", b, curr.prevHash)
		}
	}
}

func printBlock(b *block) {
	fmt.Println(b.blockData)
	fmt.Println(b.prevHash)
	fmt.Println(b)
	fmt.Println(b.link)
}

func printAllBlocks() {
	curr := head
	for curr != nil {
		printBlock(curr)
		curr = curr.link
	}
}

func main() {
	/*addBlock(20)
	addBlock(40)
	addBlock(60)
	verifyChain()
	//h := head.link
	//fmt.Println(head.link.blockData)
	//fmt.Println(h.blockData)
	fmt.Println("vim-go")*/
	var data, input, n, m, v int
	fmt.Println("1)addBlock\n 2)add n randoms blocks \n 3)alterNthBlock \n 4)print all blocks \n 5)verifyChain \n 6)hackChain\n ")
	for {
		fmt.Println("Choice")
		fmt.Scanln(&input)
		switch input {
		case 1:
			fmt.Println("enter data")
			fmt.Scanln(&data)
			addBlock(data)
			break
		case 2:
			fmt.Println("How many blocks to enter?")
			fmt.Scanln(&n)
			for i := 0; i < n; i++ {
				addBlock(rand.Intn(1250))
			}
			break
		case 3:
			fmt.Println("Wich block to alter?")
			fmt.Scanln(&m)
			fmt.Println("Wich block to enter?")
			fmt.Scanln(&v)
			alterNthBlock(m, v)
			break
		case 4:
			printAllBlocks()
			break
		case 5:
			verifyChain()
		case 6:
			hackChain()
			break
		default:
			fmt.Println("wrong choice")
			break
		}
	}
}

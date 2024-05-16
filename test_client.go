package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	testGetBlocks()

	testMineBlock()

	testAddTransaction()

	testGetChain()
}
func testGetBlocks() {
	fmt.Println("Testing GET /blocks...")
	resp, err := http.Get("http://localhost:8080/blocks")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}
func testAddTransaction() {
	fmt.Println("Testing POST /add_transaction...")
	resp, err := http.Post("http://localhost:8080/add_transaction", "application/json", bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}
func testGetChain() {
	fmt.Println("Testing GET /chain...")
	resp, err := http.Get("http://localhost:8080/chain")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}
func testMineBlock() {
	fmt.Println("Testing POST /mine_block...")
	resp, err := http.Post("http://localhost:8080/mine_block", "application/json", bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}

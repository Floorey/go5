package main

import (
	"bytes"
	"crypto/tls"
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
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := client.Get("https://localhost:8080/blocks")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}

func testAddTransaction() {
	fmt.Println("Testing POST /add_transaction...")
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := client.Post("https://localhost:8080/add_transaction", "application/json", bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}

func testGetChain() {
	fmt.Println("Testing GET /chain...")
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := client.Get("https://localhost:8080/chain")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}

func testMineBlock() {
	fmt.Println("Testing POST /mine_block...")
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := client.Post("https://localhost:8080/mine_block", "application/json", bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}

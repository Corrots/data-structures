package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/corrots/data-structures/trie"
)

func main() {
	//filename := "trie/test/a.txt"
	filename := "trie/test/pride-and-prejudice.txt"
	readFile(filename)
}

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("file open err: %v\n", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	// Trie
	t := trie.NewTrie()
	start := time.Now()
	for scanner.Scan() {
		t.Add(scanner.Text())
	}
	since := time.Since(start).Milliseconds()
	fmt.Printf("total different words: %d\n", t.Len())
	fmt.Printf("Trie: %d ms\n", since)
}

package main

import (
	"crypto/sha256"
	"fmt"
)

func DiffCount(h1 [32]byte, h2 [32]byte) int {
	var count int = 0

	for i := range 32 {
		for j := range 8 {
			var mask byte = 1 << j

			if (h1[i] & mask) != (h2[i] & mask) {
				count++
			}
		}
	}

	return count
}

func main() {
	var (
		h1 = sha256.Sum256([]byte{0})
		h2 = sha256.Sum256([]byte{1})
	)

	fmt.Printf("h1: %x\n", h1)
	fmt.Printf("h2: %x\n", h2)

	fmt.Println("DiffCount:", DiffCount(h1, h2))
}

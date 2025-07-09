package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	ptrHashType := flag.Int("t", 256, "256 | 384 | 512")

	flag.Parse()
	fmt.Println(flag.Args())

	if len(flag.Args()) > 1 {
		panic("syntax Error")
	}

	var data = []byte(flag.Args()[0])

	var hash []byte

	switch *ptrHashType {
	case 256:
		hash256 := sha256.Sum256(data)
		hash = hash256[:]
	case 384:
		hash384 := sha512.Sum384(data)
		hash = hash384[:]
	case 512:
		hash512 := sha512.Sum512(data)
		hash = hash512[:]
	}

	fmt.Printf("%x\n", hash)
}

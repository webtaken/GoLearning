package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"log"
	"os"
)

func main() {
	var hashFormat string = "256"
	if len(os.Args) >= 2 {
		hashFormat = os.Args[1]
	}

	if hashFormat != "256" && hashFormat != "384" && hashFormat != "512" {
		log.Fatalf("No provided correct sha algorithm, choose: <256|384|512>\n")
	}

	str := ""
	fmt.Printf("Insert text: ")
	fmt.Scan(&str)

	if hashFormat == "384" {
		sha384 := sha512.Sum384([]byte(str))
		fmt.Printf("Hash SHA%[2]s\n%[1]x\n%[1]T\n", sha384, hashFormat)
		os.Exit(0)
	}
	if hashFormat == "512" {
		sha512 := sha512.Sum384([]byte(str))
		fmt.Printf("Hash SHA%[2]s\n%[1]x\n%[1]T\n", sha512, hashFormat)
		os.Exit(0)
	}
	sha256 := sha256.Sum256([]byte(str))
	fmt.Printf("Hash SHA%[2]s\n%[1]x\n%[1]T\n", sha256, hashFormat)
}

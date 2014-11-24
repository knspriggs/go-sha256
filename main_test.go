package main

import (
	"crypto/sha256"
	"fmt"
	"testing"
	"time"
)

func TestEmptyString(t *testing.T) {
	fmt.Println("Kristian's Library:")

	toHash := ""

	start := time.Now()
	result := Hash(toHash)
	elapsed := time.Since(start)

	fmt.Printf("(1): ")
	PrintHash(result)
	fmt.Println("took ", elapsed, "\n")
}

func TestStringLessThan64(t *testing.T) {
	toHash := "a"

	start := time.Now()
	result := Hash(toHash)
	elapsed := time.Since(start)

	fmt.Printf("(2): ")
	//PrintHash(result)
	fmt.Printf("%x\n", result)
	fmt.Println("took ", elapsed, "\n")
}

/*func TestStringGreaterThan64(t *testing.T) {
	toHash := "this is a super long string that needs to break my program into using two seperate chunks for better testing, make sense?"

	start := time.Now()
	result := Hash(toHash)
	elapsed := time.Since(start)

	fmt.Printf("(3): ")
	PrintHash(result)
	fmt.Println("took ", elapsed, "\n")
}*/

func TestEmptyStringLib(t *testing.T) {
	fmt.Println("Standard Library:")
	start := time.Now()
	hasher := sha256.New()
	hasher.Write([]byte(""))
	result := hasher.Sum(nil)
	elapsed := time.Since(start)
	fmt.Printf("(1): %x\n", result)
	fmt.Println("took ", elapsed, "\n")
}

func TestStringLessThan64Lib(t *testing.T) {
	start := time.Now()
	hasher := sha256.New()
	hasher.Write([]byte("a"))
	result := hasher.Sum(nil)
	elapsed := time.Since(start)
	fmt.Printf("(2): %x\n", result)
	fmt.Println("took ", elapsed, "\n")
}

/*func TestStringGreaterThan64Lib(t *testing.T) {
	toHash := "this is a super long string that needs to break my program into using two seperate chunks for better testing, make sense?"

	start := time.Now()
	hasher := sha256.New()
	hasher.Write([]byte(toHash))
	result := hasher.Sum(nil)
	elapsed := time.Since(start)
	fmt.Printf("(3): %x\n", result)
	fmt.Println("took ", elapsed, "\n")
}*/

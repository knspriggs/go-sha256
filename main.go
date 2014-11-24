package main

import (
	"fmt"
	"time"
)

const (
	h0                 = 0x6a09e667
	h1                 = 0xbb67ae85
	h2                 = 0x3c6ef372
	h3                 = 0xa54ff53a
	h4                 = 0x510e527f
	h5                 = 0x9b05688c
	h6                 = 0x1f83d9ab
	h7                 = 0x5be0cd19
	BYTE_SIZE          = 8
	CHUNK_SIZE         = 512
	SIZE_BEFORE_LENGTH = 448
)

var initValues [8]uint32

var hashValueArray []chan uint32
var comm chan bool

var k = [64]uint32{
	0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5, 0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5,
	0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3, 0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
	0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc, 0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da,
	0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7, 0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967,
	0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13, 0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
	0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3, 0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070,
	0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5, 0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3,
	0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208, 0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2,
}

// create message array
func preprocessing(msg []byte) []byte {

	msg_len := len(msg) * BYTE_SIZE
	len := uint64(len(msg))

	//msg_len := byte(len(msg))
	msg = append(msg, (0x80 >> 7))

	for i := msg_len; i%CHUNK_SIZE != SIZE_BEFORE_LENGTH; i += BYTE_SIZE {
		msg = append(msg, 0x00)
	}

	len <<= 3
	for i := uint(0); i < 8; i++ {
		msg = append(msg, byte(len>>(56-8*i)))
	}

	fmt.Printf("Msg: %x\n", msg)
	return msg
}

func delegateChunks(message []byte) []uint32 {
	num_chunks := len(message) / 64
	fmt.Println("Num chunks: ", num_chunks)
	comm = make(chan bool, num_chunks)
	hashValueArray = make([]chan uint32, 8)
	for i := 0; i < 8; i++ {
		hashValueArray[i] = make(chan uint32, 10)
	}

	//go and process each chunk concurrently
	for i := 0; i < num_chunks; i++ {
		go processChunk(message[i*64:64+i*64], i)
	}

	//wait for all chunk processing functions to finish
	for i := 0; i < num_chunks; i++ {
		<-comm
	}

	//close the channels
	for i := 0; i < 8; i++ {
		close(hashValueArray[i])
	}
	return combineValues()
}

// process each chunk
func processChunk(chunk []byte, n int) {
	var w [64]uint32
	for i := 0; i < 16; i++ {
		w[i] = uint32(chunk[i])
	}

	for i := 16; i < 64; i++ {
		v1 := w[i-2]
		t1 := (v1>>17 | v1<<(32-17)) ^ (v1>>19 | v1<<(32-19)) ^ (v1 >> 10)
		v2 := w[i-15]
		t2 := (v2>>7 | v2<<(32-7)) ^ (v2>>18 | v2<<(32-18)) ^ (v2 >> 3)
		w[i] = t1 + w[i-7] + t2 + w[i-16]
	}

	a := initValues[0]
	b := initValues[1]
	c := initValues[2]
	d := initValues[3]
	e := initValues[4]
	f := initValues[5]
	g := initValues[6]
	h := initValues[7]

	for i := 0; i < 64; i++ {
		s0 := h + ((e>>6 | e<<(32-6)) ^ (e>>11 | e<<(32-11)) ^ (e>>25 | e<<(32-25))) + ((e & f) ^ (^e & g)) + k[i] + w[i]
		s1 := ((a>>2 | a<<(32-2)) ^ (a>>13 | a<<(32-13)) ^ (a>>22 | a<<(32-22))) + ((a & b) ^ (a & c) ^ (b & c))

		h = g
		g = f
		f = e
		e = d + s0
		d = c
		c = b
		b = a
		a = s0 + s1
	}

	hashValueArray[0] <- a
	hashValueArray[1] <- b
	hashValueArray[2] <- c
	hashValueArray[3] <- d
	hashValueArray[4] <- e
	hashValueArray[5] <- f
	hashValueArray[6] <- g
	hashValueArray[7] <- h
	comm <- true
}

func combineValues() []uint32 {
	res := [8]uint32{0, 0, 0, 0, 0, 0, 0, 0}
	var result []uint32
	for i := 0; i < 8; i++ {
		for v := range hashValueArray[i] {
			res[i] += v
		}
	}
	for i := 0; i < 8; i++ {
		result = append(result, res[i])
	}
	return result
}

func setup() {
	initValues[0], initValues[1], initValues[2], initValues[3], initValues[4], initValues[5], initValues[6], initValues[7] = h0, h1, h2, h3, h4, h5, h6, h7
}

func Hash(msg string) []uint32 {
	setup()
	msg_p := preprocessing([]byte(msg))
	return delegateChunks(msg_p)
}

func PrintHash(ar []uint32) {
	for i := 0; i < len(ar); i++ {
		fmt.Printf("%x", ar[i])
	}
	fmt.Printf("\n")
}

func main() {
	//toHash := "this is a super long string that needs to break my program into using two seperate chunks for better testing, make sense?"
	toHash := ""

	start := time.Now()
	result := Hash(toHash)
	elapsed := time.Since(start)

	fmt.Printf("Hash: %x\n", result)

	//PrintHash(result)
	fmt.Println("took ", elapsed)
}

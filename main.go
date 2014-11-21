package main

import (
	//"encoding/hex"
	"fmt"
	"strings"
)

const (
	h0   = 0x6a09e667
	h1   = 0xbb67ae85
	h2   = 0x3c6ef372
	h3   = 0xa54ff53a
	h4   = 0x510e527f
	h5   = 0x9b05688c
	h6   = 0x1f83d9ab
	h7   = 0x5be0cd19
	zero = 0x0
	one  = 0x1
)

var hashValueArray []chan uint32

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

func rightRotate(b []uint32, x int) [8]uint32 {
	len := len(b)
	fmt.Println(len)
	var tmp [8]uint32
	for i := 0; i < len; i++ {
		fmt.Printf("%d: %d -> %d: %d", i, b[i], i, b[(i+x)%len])
		tmp[i] = b[(i+x)%len]
	}
	return tmp
}

// break the message into chunks
func preprocessing(message string) []uint32 {
	msg_s := strings.Split(message, "")
	var msg []uint32
	for c := range msg_s {
		msg = append(msg, uint32(c))
	}
	msg_len := len(msg)
	msg = append(msg, one)

	for (len(msg) % 64) < 56 {
		msg = append(msg, zero)
	}
	for i := uint(0); i < 8; i++ {
		msg = append(msg, uint32(msg_len>>(56-8*i)))
	}
	fmt.Printf("%v  -  %v\n", msg, len(msg))
	return msg
}

func delegateChunks(message []uint32) {
	num_chunks := len(message) / 64
	//comm := make([]chan bool, num_chunks)
	hashValueArray = make([]chan uint32, 8)
	for i := 0; i < 8; i++ {
		hashValueArray[i] = make(chan uint32, 10)
	}
	for i := 0; i < num_chunks; i++ {
		go processChunk(message[i*64:64+i*64], i)
	}
}

// process each chunk
func processChunk(chunk []uint32, n int) {
	fmt.Printf("Working on chunk %d\n", n)
	fmt.Printf("Chunk data: %v\n", chunk)
	var w [64]uint32
	for i := 0; i < 16; i++ {
		w[i] = chunk[i]
	}

	for i := 16; i < 64; i++ {
		s0 := (w[i-15] >> 7) ^ (w[i-15] >> 18) ^ (w[i-15] >> 3)
		s1 := (w[i-2] >> 17) ^ (w[i-2] >> 19) ^ (w[i-2] >> 10)
		w[i] = w[i-16] + s0 + w[i-7] + s1
	}

	a := h0
	b := h1
	c := h2
	d := h3
	e := h4
	f := h5
	g := h6
	h := h7

	for i := 0; i < 64; i++ {
		s1 := (e >> 6) ^ (e >> 11) ^ (e >> 25)
		ch := (e & f) ^ (g &^ e)
		temp1 := uint32(h) + uint32(s1) + uint32(ch) + k[i] + uint32(w[i])
		s0 := (a >> 2) ^ (a >> 13) ^ (a >> 22)
		maj := (a & b) ^ (a & c) ^ (b & c)
		temp2 := s0 + maj

		h = g
		g = f
		f = e
		e = d + int(temp1)
		d = c
		c = b
		b = a
		a = int(temp1) + temp2
	}
	fmt.Printf("Adding values to channels for chunk %d\n", n)
	hashValueArray[0] <- uint32(a)
	hashValueArray[1] <- uint32(b)
	hashValueArray[2] <- uint32(c)
	hashValueArray[3] <- uint32(d)
	hashValueArray[4] <- uint32(e)
	hashValueArray[5] <- uint32(f)
	hashValueArray[6] <- uint32(g)
	hashValueArray[7] <- uint32(h)
}

/*func sum(channel chan int) byte {
	res := 0
	for val := range channel {
		res += val
	}
	return byte(res)
}*/

func combineValues() []uint32 {
	res := [8]uint32{0, 0, 0, 0, 0, 0, 0, 0}
	var result []uint32
	for i := 0; i < 8; i++ {
		for v := range hashValueArray[i] {
			res[i] += v //<-hashValueArray[i]
		}
	}
	for i := 0; i < 8; i++ {
		result = append(result, res[i])
	}
	return result
}

func main() {
	fmt.Println("----Start----")

	toHash := "this"
	fmt.Println("In length: ", len(toHash))
	msg := preprocessing(toHash)
	delegateChunks(msg)

	result := combineValues()
	fmt.Printf("%x\n", result)
	//fmt.Println(hex.EncodeToString(result[0]))
}

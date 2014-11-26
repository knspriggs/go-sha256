package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"strings"
	"time"
)

const (
	h0 = 0x6a09e667
	h1 = 0xbb67ae85
	h2 = 0x3c6ef372
	h3 = 0xa54ff53a
	h4 = 0x510e527f
	h5 = 0x9b05688c
	h6 = 0x1f83d9ab
	h7 = 0x5be0cd19
)

var currentHashValues [8]uint32

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
	msg_len := uint64(len(msg))
	msg = append(msg, 0x80)

	for len(msg)%64 < 56 {
		msg = append(msg, 0x00)
	}

	msg_len *= 8

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, msg_len)
	for _, b := range buf.Bytes() {
		msg = append(msg, b)
	}
	return msg
}

func delegateChunks(message []byte) [8]uint32 {
	num_chunks := len(message) / 64

	//process each chunk
	for i := 0; i < num_chunks; i++ {
		processChunk(message[i*64:64+i*64], i)
	}
	return currentHashValues
}

// process each chunk
func processChunk(chunk []byte, n int) {
	var w [64]uint32
	for i := 0; i < 16; i++ {
		j := i * 4
		w[i] = uint32(chunk[j])<<24 | uint32(chunk[j+1])<<16 | uint32(chunk[j+2])<<8 | uint32(chunk[j+3])
	}

	for i := 16; i < 64; i++ {
		v1 := w[i-2]
		t1 := (v1>>17 | v1<<(32-17)) ^ (v1>>19 | v1<<(32-19)) ^ (v1 >> 10)
		v2 := w[i-15]
		t2 := (v2>>7 | v2<<(32-7)) ^ (v2>>18 | v2<<(32-18)) ^ (v2 >> 3)
		w[i] = t1 + w[i-7] + t2 + w[i-16]
	}

	a := currentHashValues[0]
	b := currentHashValues[1]
	c := currentHashValues[2]
	d := currentHashValues[3]
	e := currentHashValues[4]
	f := currentHashValues[5]
	g := currentHashValues[6]
	h := currentHashValues[7]

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

	currentHashValues[0] = currentHashValues[0] + a
	currentHashValues[1] = currentHashValues[1] + b
	currentHashValues[2] = currentHashValues[2] + c
	currentHashValues[3] = currentHashValues[3] + d
	currentHashValues[4] = currentHashValues[4] + e
	currentHashValues[5] = currentHashValues[5] + f
	currentHashValues[6] = currentHashValues[6] + g
	currentHashValues[7] = currentHashValues[7] + h
}

func setup() {
	currentHashValues[0], currentHashValues[1], currentHashValues[2], currentHashValues[3] = h0, h1, h2, h3
	currentHashValues[4], currentHashValues[5], currentHashValues[6], currentHashValues[7] = h4, h5, h6, h7
}

func Hash(msg string) [8]uint32 {
	setup()
	msg_p := preprocessing([]byte(msg))
	return delegateChunks(msg_p)
}

func StringValue(ar [8]uint32) string {
	var ar_s []string
	for _, v := range ar {
		if len(fmt.Sprintf("%x", v)) != 8 {
			ar_s = append(ar_s, "0")
		}
		ar_s = append(ar_s, fmt.Sprintf("%x", v))
	}
	return strings.Join(ar_s, "")
}

type sha256Test struct {
	out string
	in  string
}

var golden = []sha256Test{
	{"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", ""},
	{"ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb", "a"},
	{"fb8e20fc2e4c3f248c60c39bd652f3c1347298bb977b8b4d5903b85055620603", "ab"},
	{"ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad", "abc"},
	{"88d4266fd4e6338d13b845fcf289579d209c897823b9217da3e161936f031589", "abcd"},
	{"36bbe50ed96841d10443bcb670d6554f0a34b761be67ec9c4a8ad2c0c44ca42c", "abcde"},
	{"bef57ec7f53a6d40beb640a780a639c83bc29ac8a9816f1fc6c5c6dcd93c4721", "abcdef"},
	{"7d1a54127b222502f5b79b5fb0803061152a44f92b37e23c6527baf665d4da9a", "abcdefg"},
	{"9c56cc51b374c3ba189210d5b6d4bf57790d351c96c47c02190ecf1e430635ab", "abcdefgh"},
	{"ae7a702a9509039ddbf29f0765e70d0001177914b86459284dab8b348c2dce3f", "I wouldn't marry him with a ten foot pole."},
	{"7102cfd76e2e324889eece5d6c41921b1e142a4ac5a2692be78803097f6a48d8", "Nepal premier won't resign."},
	{"23b1018cd81db1d67983c5f7417c44da9deb582459e378d7a068552ea649dc9f", "For every action there is an equal and opposite government program."},
	{"8c87deb65505c3993eb24b7a150c4155e82eee6960cf0c3a8114ff736d69cad5", "There is no reason for any individual to have a computer in their home. -Ken Olsen, 1977"},
	{"bfb0a67a19cdec3646498b2e0f751bddc41bba4b7f30081b0b932aad214d16d7", "It's a tiny change to the code and not completely disgusting. - Bob Manchek"},
	{"7f9a0b9bf56332e19f5a0ec1ad9c1425a153da1c624868fda44561d6b74daf36", "size:  a.out:  bad magic"},
	{"395585ce30617b62c80b93e8208ce866d4edc811a177fdb4b82d3911d8696423", "The fugacity of a constituent in a mixture of gases at a given temperature is proportional to its mole fraction.  Lewis-Randall Rule"},
}

func main() {
	fmt.Println("----- Correctness Tests ------")
	for _, t := range golden {
		fmt.Printf("In: %s\n", t.in)
		start := time.Now()
		res := Hash(t.in)
		elapsed := time.Since(start)
		res_s := StringValue(res)
		if res_s == t.out {
			fmt.Printf("%s matches expected\n", res_s)
		} else {
			fmt.Printf("%s does not match %s\n", res_s, t.out)
		}
		fmt.Println("Took: ", elapsed)
		fmt.Println("-----")
	}
	fmt.Println("-------------------------------")

	fmt.Println("--------- Speed Tests ---------")
	for _, t := range golden {
		fmt.Printf("In: %s\n", t.in)
		start := time.Now()
		Hash(t.in)
		elapsed1 := time.Since(start)
		fmt.Println("My library:\t\t", elapsed1)

		start = time.Now()
		hasher := sha256.New()
		hasher.Write([]byte(t.in))
		hasher.Sum(nil)
		elapsed2 := time.Since(start)
		fmt.Println("Standard Library:\t", elapsed2)
		fmt.Println("-----")
	}
	fmt.Println("-------------------------------")
}

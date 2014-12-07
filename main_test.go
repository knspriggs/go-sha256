package main

import (
  "testing"
  "crypto/sha256"
  "time"
  )

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

func TestCorrectness(t *testing.T) {
  for _, g := range golden {
		res := Hash(g.in)
		res_s := StringValue(res)
		if res_s != g.out {
			t.Fatalf("%s does not matche expected", res_s)
		}
	}
}

func TestSpeed(t *testing.T) {
  for _, g := range golden {
    start := time.Now()
    Hash(g.in)
    elapsed1 := time.Since(start)

    start = time.Now()
    hasher := sha256.New()
    hasher.Write([]byte(g.in))
    hasher.Sum(nil)
    elapsed2 := time.Since(start)
    if elapsed1 >= (10 * elapsed2) {
      t.Fatalf("Too slow\n")
    }
  }
}

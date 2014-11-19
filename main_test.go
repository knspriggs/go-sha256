package main

import (
	"fmt"
	"testing"
)

func TestRightRotate(t *testing.T) {
	b := []byte{10, 34, 41, 51}
	fmt.Println(rightRotate(b, 2))
	/*if (rightRotate(b, 2) != [41, 51, 10, 34]) {
	  t.Error("Not correct")
	}*/
}

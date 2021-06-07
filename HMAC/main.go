package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	a := getCode("aaa@bnjne.com")
	fmt.Println(a)
	a = getCode("nniekne#jngie@jn.com")
	fmt.Println(a)
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("mykey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

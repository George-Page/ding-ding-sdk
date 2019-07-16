package dingTalk

import (
	cr "crypto/rand"
	mr "math/rand"
	"time"
	"bytes"
)


func Encode(cipherText []byte, blockSize int) []byte {
	if blockSize == 0 {
		blockSize = PKCS7EncoderBlockSize
	}
	textLen := len(cipherText)
	padding := blockSize - (textLen) % blockSize
	if padding == 0 {
		padding = blockSize
	}
	return append(cipherText, bytes.Repeat([]byte{byte(padding)}, padding)...)
}

func Decode(cipherText []byte) []byte {
	padding := int(cipherText[len(cipherText) -1])
	if padding < 1 {
		padding = 0
	}
	return cipherText[0:len(cipherText) - padding]
}

func GetRandomString(n int, alphabets ...byte) string {
	if n == 0 {
		n = DefaultRandomNum
	}

	bts := make([]byte, n)
	randBoolean := false
	if num, err := cr.Read(bts); num != n || err != nil {
		mr.Seed(time.Now().UnixNano())
		randBoolean = true
	}
	for i, b := range bts {
		if len(alphabets) == 0 {
			if randBoolean {
				bts[i] = AlphabetsPool[mr.Intn(len(AlphabetsPool))]
			}else{
				bts[i] = AlphabetsPool[b%byte(len(AlphabetsPool))]
			}
		}else {
			if randBoolean {
				bts[i] = alphabets[mr.Intn(len(alphabets))]
			}else{
				bts[i] = alphabets[b%byte(len(alphabets))]
			}
		}
	}
	return string(bts)
}



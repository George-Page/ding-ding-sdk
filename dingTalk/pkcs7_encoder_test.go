package dingTalk

import (
	"testing"
	"github.com/micro/go-log"
)

func TestEncode(t *testing.T) {
	cipherText := []byte("x13579y")
	en := Encode(cipherText, 32)
	log.Logf("original string:%v", string(cipherText))
	log.Logf("original bytes:%v", cipherText)
	log.Logf("encode bytes:%v", en)
}

func TestDecode(t *testing.T) {
	cipherText := []byte("x13579y")
	msg := Encode(cipherText, 32)
	de := Decode(msg)
	log.Logf("original string:%v", string(msg))
	log.Logf("decode bytes:%v", de)
}

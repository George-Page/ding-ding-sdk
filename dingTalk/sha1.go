package dingTalk

import (
	"crypto/sha1"
	"fmt"
)

func Sha1Sign(str string) string {
	s1 := sha1.New()
	s1.Write([]byte(str))
	return fmt.Sprintf("%x", s1.Sum(nil))
}

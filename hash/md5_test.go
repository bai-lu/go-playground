package hash

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	origin := []byte("hello world")
	a := md5Sum(origin)
	b := hasherSum(origin)
	if fmt.Sprintf("%x\n", a) != fmt.Sprintf("%x\n", b) {
		t.Errorf(fmt.Sprintf("%x != %x", a, b))
	}
}

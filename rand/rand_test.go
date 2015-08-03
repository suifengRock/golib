package rand

import (
	"math/rand"
	"testing"
)

func TestStringRand(t *testing.T) {
	res := StringRand(Lowercase, 0)
	if res != "" {
		t.Error("StringRand req=0 is err")
	}
	req_num := 16
	res = StringRand(Capital, req_num)
	if len(res) != req_num {
		t.Error("StringRand is err")
	}
}

func TestSysRand(t *testing.T) {
	str, err := SysRand(32)
	if err != nil {
		t.Error(err)
	}
	t.Log(str)
}

func BenchmarkStringRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoLimit(32)
	}
}

func BenchmarkFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SysRand(32)
	}
}

func BenchmarkFloat32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Float64()
	}
}

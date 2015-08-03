package rand

import (
	"crypto/rand"
	"fmt"
)

// 用系统 /dev/urandom 生成伪随机数
func SysRand(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	str := fmt.Sprintf("%x", b[:])
	return str, nil
}

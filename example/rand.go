package main

import (
	"fmt"
	"golib/rand"
)

func main() {
	fmt.Println("")
	fmt.Println(rand.StringRand(rand.Capital+rand.SpeStr, 32), "  32")

	fmt.Println(rand.NoLimit(32), "  32")

	fmt.Println(rand.CapitalOnly(32), "  32")

	fmt.Println(rand.LowercaseOnly(32), "  32")

	fmt.Println(rand.NumberOnly(32), "  32")

	fmt.Println(rand.SpeStrOnly(32), "  32")

	fmt.Println(rand.NoSpeStr(32), "  32")
	fmt.Println("")
}

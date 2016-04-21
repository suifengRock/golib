package ticker

import (
	"testing"
	"time"
)

func SayTest() {
	time.Sleep(1 * time.Minute)
	// fmt.Println("test ticker.....")
}

func BenchmarkTickerHandle(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go TickerHandle(SayTest, 5*time.Second)
	}
}

func BenchmarkTickerHandle2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go TickerHandle(SayTest, 2*time.Second)
	}
}

func BenchmarkTickerHandle3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go TickerHandle(SayTest, 1*time.Second)
	}
}

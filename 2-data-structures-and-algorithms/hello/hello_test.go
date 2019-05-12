package hello

import (
	"fmt"
	"testing"
)

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("Hello High Performance Go") // Print with no formatting
	}
}

func BenchmarkGoodbye(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("Goodbye High Performance Go") // Print with no formatting
	}
}
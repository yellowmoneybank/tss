package benchmarks

import (
	"testing"

	"moritzm-mueller.de/tss/pkg/shamir"
)

// Splits single serect 1 share 1 threshold
func BenchmarkSplitShares5(b *testing.B) {
	size := 100
	secret := make([]byte, size)
	initialize(secret)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		numberOfShares := 5
		threshold := 3

		_, err := shamir.SplitSecret(secret, numberOfShares, uint8(threshold))
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkSplitShares10(b *testing.B) {
	size := 100
	secret := make([]byte, size)
	initialize(secret)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		numberOfShares := 10
		threshold := 3

		_, err := shamir.SplitSecret(secret, numberOfShares, uint8(threshold))
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkSplitShares100(b *testing.B) {
	size := 100
	secret := make([]byte, size)
	initialize(secret)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		numberOfShares := 100
		threshold := 3

		_, err := shamir.SplitSecret(secret, numberOfShares, uint8(threshold))
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkSplitShares300(b *testing.B) {
	size := 100
	secret := make([]byte, size)
	initialize(secret)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		numberOfShares := 300
		threshold := 3

		_, err := shamir.SplitSecret(secret, numberOfShares, uint8(threshold))
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkSplitShares500(b *testing.B) {
	size := 100
	secret := make([]byte, size)
	initialize(secret)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		numberOfShares := 500
		threshold := 3

		_, err := shamir.SplitSecret(secret, numberOfShares, uint8(threshold))
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkSplitShares800(b *testing.B) {
	size := 100
	secret := make([]byte, size)
	initialize(secret)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		numberOfShares := 800
		threshold := 3

		_, err := shamir.SplitSecret(secret, numberOfShares, uint8(threshold))
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkSplitShares1000(b *testing.B) {
	size := 100
	secret := make([]byte, size)
	initialize(secret)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		numberOfShares := 1000
		threshold := 3

		_, err := shamir.SplitSecret(secret, numberOfShares, uint8(threshold))
		if err != nil {
			println(err)
		}
	}
}

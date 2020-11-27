package benchmarks

import (
	"testing"

	"moritzm-mueller.de/tss/pkg/shamir"
)

func BenchmarkSplitSecretThreshold1(b *testing.B) {
	size := 100
	secret := make([]byte, size)
	initialize(secret)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		numberOfShares := 100
		threshold := 1

		_, err := shamir.SplitSecret(secret, numberOfShares, uint8(threshold))
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkSplitSecretThreshold10(b *testing.B) {
	size := 100
	secret := make([]byte, size)
	initialize(secret)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		numberOfShares := 100
		threshold := 10

		_, err := shamir.SplitSecret(secret, numberOfShares, uint8(threshold))
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkSplitSecretThreshold20(b *testing.B) {
	size := 100
	secret := make([]byte, size)
	initialize(secret)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		numberOfShares := 100
		threshold := 20

		_, err := shamir.SplitSecret(secret, numberOfShares, uint8(threshold))
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkSplitSecretThreshold40(b *testing.B) {
	size := 100
	secret := make([]byte, size)
	initialize(secret)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		numberOfShares := 100
		threshold := 40

		_, err := shamir.SplitSecret(secret, numberOfShares, uint8(threshold))
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkSplitSecretThreshold60(b *testing.B) {
	size := 100
	secret := make([]byte, size)
	initialize(secret)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		numberOfShares := 100
		threshold := 60

		_, err := shamir.SplitSecret(secret, numberOfShares, uint8(threshold))
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkSplitSecretThreshold80(b *testing.B) {
	size := 100
	secret := make([]byte, size)
	initialize(secret)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		numberOfShares := 100
		threshold := 80

		_, err := shamir.SplitSecret(secret, numberOfShares, uint8(threshold))
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkSplitSecretThreshold100(b *testing.B) {
	size := 100
	secret := make([]byte, size)
	initialize(secret)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		numberOfShares := 100
		threshold := 100

		_, err := shamir.SplitSecret(secret, numberOfShares, uint8(threshold))
		if err != nil {
			println(err)
		}
	}
}

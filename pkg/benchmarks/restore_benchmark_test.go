package benchmarks

import (
	"testing"

	"moritzm-mueller.de/tss/pkg/shamir"
)

func BenchmarkRestoreThreshold1(b *testing.B) {
	threshold := 2
	shares := 20

	size := 100
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := shamir.Reconstruct(splits)
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkRestoreThreshold5(b *testing.B) {
	threshold := 5
	shares := 20

	size := 100
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := shamir.Reconstruct(splits)
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkRestoreThreshold10(b *testing.B) {
	threshold := 10
	shares := 20

	size := 100
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := shamir.Reconstruct(splits)
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkRestoreThreshold15(b *testing.B) {
	threshold := 15
	shares := 20

	size := 100
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := shamir.Reconstruct(splits)
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkRestoreShare1(b *testing.B) {
	threshold := 1
	shares := 1

	size := 100
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := shamir.Reconstruct(splits)
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkRestoreShare10(b *testing.B) {
	threshold := 1
	shares := 10

	size := 100
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := shamir.Reconstruct(splits)
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkRestoreShare20(b *testing.B) {
	threshold := 1
	shares := 20

	size := 100
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := shamir.Reconstruct(splits)
		if err != nil {
			println(err)
		}
	}
}

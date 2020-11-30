package benchmarks

import (
	"testing"

	"moritzm-mueller.de/tss/pkg/shamir"
)

func BenchmarkRestoreThreshold1(b *testing.B) {
	threshold := 1
	shares := 1

	size := 12000
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		shamir.Reconstruct(splits[0:threshold])
	}
}

func BenchmarkRestoreThreshold2(b *testing.B) {
	threshold := 2
	shares := 2

	size := 12000
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		shamir.Reconstruct(splits[0:threshold])
	}
}

func BenchmarkRestoreThreshold3(b *testing.B) {
	threshold := 3
	shares := 3

	size := 12000
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		shamir.Reconstruct(splits[0:threshold])
	}
}

func BenchmarkRestoreThreshold4(b *testing.B) {
	threshold := 4
	shares := 4

	size := 12000
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		shamir.Reconstruct(splits[0:threshold])
	}
}

func BenchmarkRestoreThreshold5(b *testing.B) {
	threshold := 5
	shares := 5

	size := 12000
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		shamir.Reconstruct(splits[0:threshold])
	}
}

func BenchmarkRestoreThreshold6(b *testing.B) {
	threshold := 6
	shares := 6

	size := 12000
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		shamir.Reconstruct(splits[0:threshold])
	}
}

func BenchmarkRestoreThreshold7(b *testing.B) {
	threshold := 7
	shares := 7

	size := 12000
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		shamir.Reconstruct(splits[0:threshold])
	}
}

func BenchmarkRestoreThreshold8(b *testing.B) {
	threshold := 8
	shares := 8

	size := 12000
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		shamir.Reconstruct(splits[0:threshold])
	}
}

func BenchmarkRestoreThreshold9(b *testing.B) {
	threshold := 9
	shares := 9

	size := 12000
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		shamir.Reconstruct(splits[0:threshold])
	}
}

func BenchmarkRestoreThreshold10(b *testing.B) {
	threshold := 10
	shares := 10

	size := 12000
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		shamir.Reconstruct(splits[0:threshold])
	}
}

func BenchmarkRestoreThreshold11(b *testing.B) {
	threshold := 11
	shares := 11

	size := 12000
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		shamir.Reconstruct(splits[0:threshold])
	}
}

func BenchmarkRestoreThreshold12(b *testing.B) {
	threshold := 12
	shares := 12

	size := 12000
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		shamir.Reconstruct(splits[0:threshold])
	}
}

func BenchmarkRestoreThreshold13(b *testing.B) {
	threshold := 13
	shares := 13

	size := 12000
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		shamir.Reconstruct(splits[0:threshold])
	}
}

func BenchmarkRestoreThreshold14(b *testing.B) {
	threshold := 14
	shares := 14

	size := 12000
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		shamir.Reconstruct(splits[0:threshold])
	}
}

func BenchmarkRestoreShare1KB(b *testing.B) {
	threshold := 3
	shares := 3

	size := 1000
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		shamir.Reconstruct(splits[0:threshold])
	}
}

func BenchmarkRestoreShare100KB(b *testing.B) {
	threshold := 3
	shares := 3

	size := 100000
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		shamir.Reconstruct(splits[0:threshold])
	}
}

func BenchmarkRestoreShare500KB(b *testing.B) {
	threshold := 3
	shares := 3

	size := 500000
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		shamir.Reconstruct(splits[0:threshold])
	}
}

func BenchmarkRestoreShare1MB(b *testing.B) {
	threshold := 3
	shares := 3

	size := 1000000
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		shamir.Reconstruct(splits[0:threshold])
	}
}

func BenchmarkRestoreShare100MB(b *testing.B) {
	threshold := 3
	shares := 3

	size := 100000000
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		shamir.Reconstruct(splits[0:threshold])
	}
}

func BenchmarkRestoreShare500MB(b *testing.B) {
	threshold := 3
	shares := 3

	size := 500000000
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares, uint8(threshold))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		shamir.Reconstruct(splits[0:threshold])
	}
}

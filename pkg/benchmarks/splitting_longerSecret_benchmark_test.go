package benchmarks

import (
	"io/ioutil"
	"testing"

	"moritzm-mueller.de/tss/pkg/shamir"
)

func BenchmarkSplitSecret1MB(b *testing.B) {
	secret, _ := ioutil.ReadFile("./1MB.txt")

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

func BenchmarkSplitSecretBytes1(b *testing.B) {
	size := 1
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

func BenchmarkSplitSecretBytes10(b *testing.B) {
	size := 10
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

func BenchmarkSplitSecretBytes100(b *testing.B) {
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

func BenchmarkSplitSecretBytes500(b *testing.B) {
	size := 500
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

func BenchmarkSplitSecretBytes1000(b *testing.B) {
	size := 1000
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

func initialize(s []byte) {
	for i := range s {
		s[i] = 1
	}
}

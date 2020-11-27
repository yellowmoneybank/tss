package benchmarks

import (
	"testing"

	"moritzm-mueller.de/tss/pkg/redistribute"
	"moritzm-mueller.de/tss/pkg/secretSharing"

	"moritzm-mueller.de/tss/pkg/shamir"
)

func BenchmarkRedistBase(b *testing.B) {
	size := 100
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, 5, 2)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := redistribute.RedistributeShare(splits[0], 5, 3)
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkRedistReconstruct(b *testing.B) {
	threshold_old := 2
	shares_old := 5

	threshold_new := 5
	shares_new := 7

	size := 100
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares_old, uint8(threshold_old))

	var redistShares [][]secretSharing.RedistShare

	for _, split := range splits {
		redist, err := redistribute.RedistributeShare(split, shares_new, threshold_new)
		if err != nil {
			println(err)
		}

		redistShares = append(redistShares, redist)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := redistribute.RedistSharesToShareholders(redistShares)
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkRedistReconstructShares10(b *testing.B) {
	threshold_old := 2
	shares_old := 5

	threshold_new := 2
	shares_new := 10

	size := 100
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares_old, uint8(threshold_old))

	var redistShares [][]secretSharing.RedistShare

	for _, split := range splits {
		redist, err := redistribute.RedistributeShare(split, shares_new, threshold_new)
		if err != nil {
			println(err)
		}

		redistShares = append(redistShares, redist)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := redistribute.RedistSharesToShareholders(redistShares)
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkRedistReconstructShares100(b *testing.B) {
	threshold_old := 2
	shares_old := 5

	threshold_new := 2
	shares_new := 100

	size := 100
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares_old, uint8(threshold_old))

	var redistShares [][]secretSharing.RedistShare

	for _, split := range splits {
		redist, err := redistribute.RedistributeShare(split, shares_new, threshold_new)
		if err != nil {
			println(err)
		}

		redistShares = append(redistShares, redist)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := redistribute.RedistSharesToShareholders(redistShares)
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkRedistReconstructShares300(b *testing.B) {
	threshold_old := 2
	shares_old := 5

	threshold_new := 2
	shares_new := 300

	size := 100
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares_old, uint8(threshold_old))

	var redistShares [][]secretSharing.RedistShare

	for _, split := range splits {
		redist, err := redistribute.RedistributeShare(split, shares_new, threshold_new)
		if err != nil {
			println(err)
		}

		redistShares = append(redistShares, redist)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := redistribute.RedistSharesToShareholders(redistShares)
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkRedistReconstructShares500(b *testing.B) {
	threshold_old := 2
	shares_old := 5

	threshold_new := 2
	shares_new := 500

	size := 100
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares_old, uint8(threshold_old))

	var redistShares [][]secretSharing.RedistShare

	for _, split := range splits {
		redist, err := redistribute.RedistributeShare(split, shares_new, threshold_new)
		if err != nil {
			println(err)
		}

		redistShares = append(redistShares, redist)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := redistribute.RedistSharesToShareholders(redistShares)
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkRedistReconstructShares800(b *testing.B) {
	threshold_old := 2
	shares_old := 5

	threshold_new := 2
	shares_new := 800

	size := 100
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares_old, uint8(threshold_old))

	var redistShares [][]secretSharing.RedistShare

	for _, split := range splits {
		redist, err := redistribute.RedistributeShare(split, shares_new, threshold_new)
		if err != nil {
			println(err)
		}

		redistShares = append(redistShares, redist)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := redistribute.RedistSharesToShareholders(redistShares)
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkRedistReconstructShares1000(b *testing.B) {
	threshold_old := 2
	shares_old := 5

	threshold_new := 2
	shares_new := 1000

	size := 100
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares_old, uint8(threshold_old))

	var redistShares [][]secretSharing.RedistShare

	for _, split := range splits {
		redist, err := redistribute.RedistributeShare(split, shares_new, threshold_new)
		if err != nil {
			println(err)
		}

		redistShares = append(redistShares, redist)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := redistribute.RedistSharesToShareholders(redistShares)
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkRedistReconstructSharesThreshold1(b *testing.B) {
	threshold_old := 2
	shares_old := 5

	threshold_new := 1
	shares_new := 10

	size := 10
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares_old, uint8(threshold_old))

	var redistShares [][]secretSharing.RedistShare

	for _, split := range splits {
		redist, err := redistribute.RedistributeShare(split, shares_new, threshold_new)
		if err != nil {
			println(err)
		}

		redistShares = append(redistShares, redist)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := redistribute.RedistSharesToShareholders(redistShares)
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkRedistReconstructSharesThreshold2(b *testing.B) {
	threshold_old := 2
	shares_old := 5

	threshold_new := 2
	shares_new := 10

	size := 10
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares_old, uint8(threshold_old))

	var redistShares [][]secretSharing.RedistShare

	for _, split := range splits {
		redist, err := redistribute.RedistributeShare(split, shares_new, threshold_new)
		if err != nil {
			println(err)
		}

		redistShares = append(redistShares, redist)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := redistribute.RedistSharesToShareholders(redistShares)
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkRedistReconstructSharesThreshold3(b *testing.B) {
	threshold_old := 2
	shares_old := 5

	threshold_new := 3
	shares_new := 10

	size := 10
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares_old, uint8(threshold_old))

	var redistShares [][]secretSharing.RedistShare

	for _, split := range splits {
		redist, err := redistribute.RedistributeShare(split, shares_new, threshold_new)
		if err != nil {
			println(err)
		}

		redistShares = append(redistShares, redist)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := redistribute.RedistSharesToShareholders(redistShares)
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkRedistReconstructSharesThreshold4(b *testing.B) {
	threshold_old := 2
	shares_old := 5

	threshold_new := 4
	shares_new := 10

	size := 10
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares_old, uint8(threshold_old))

	var redistShares [][]secretSharing.RedistShare

	for _, split := range splits {
		redist, err := redistribute.RedistributeShare(split, shares_new, threshold_new)
		if err != nil {
			println(err)
		}

		redistShares = append(redistShares, redist)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := redistribute.RedistSharesToShareholders(redistShares)
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkRedistReconstructSharesThreshold5(b *testing.B) {
	threshold_old := 2
	shares_old := 5

	threshold_new := 5
	shares_new := 10

	size := 10
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares_old, uint8(threshold_old))

	var redistShares [][]secretSharing.RedistShare

	for _, split := range splits {
		redist, err := redistribute.RedistributeShare(split, shares_new, threshold_new)
		if err != nil {
			println(err)
		}

		redistShares = append(redistShares, redist)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := redistribute.RedistSharesToShareholders(redistShares)
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkRedistReconstructSharesThreshold6(b *testing.B) {
	threshold_old := 2
	shares_old := 5

	threshold_new := 6
	shares_new := 10

	size := 10
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares_old, uint8(threshold_old))

	var redistShares [][]secretSharing.RedistShare

	for _, split := range splits {
		redist, err := redistribute.RedistributeShare(split, shares_new, threshold_new)
		if err != nil {
			println(err)
		}

		redistShares = append(redistShares, redist)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := redistribute.RedistSharesToShareholders(redistShares)
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkRedistReconstructSharesThreshold7(b *testing.B) {
	threshold_old := 2
	shares_old := 5

	threshold_new := 7
	shares_new := 10

	size := 10
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares_old, uint8(threshold_old))

	var redistShares [][]secretSharing.RedistShare

	for _, split := range splits {
		redist, err := redistribute.RedistributeShare(split, shares_new, threshold_new)
		if err != nil {
			println(err)
		}

		redistShares = append(redistShares, redist)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := redistribute.RedistSharesToShareholders(redistShares)
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkRedistReconstructSharesThreshold8(b *testing.B) {
	threshold_old := 2
	shares_old := 5

	threshold_new := 8
	shares_new := 10

	size := 10
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares_old, uint8(threshold_old))

	var redistShares [][]secretSharing.RedistShare

	for _, split := range splits {
		redist, err := redistribute.RedistributeShare(split, shares_new, threshold_new)
		if err != nil {
			println(err)
		}

		redistShares = append(redistShares, redist)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := redistribute.RedistSharesToShareholders(redistShares)
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkRedistReconstructSharesThreshold9(b *testing.B) {
	threshold_old := 2
	shares_old := 5

	threshold_new := 9
	shares_new := 10

	size := 10
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares_old, uint8(threshold_old))

	var redistShares [][]secretSharing.RedistShare

	for _, split := range splits {
		redist, err := redistribute.RedistributeShare(split, shares_new, threshold_new)
		if err != nil {
			println(err)
		}

		redistShares = append(redistShares, redist)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := redistribute.RedistSharesToShareholders(redistShares)
		if err != nil {
			println(err)
		}
	}
}

func BenchmarkRedistReconstructSharesThreshold10(b *testing.B) {
	threshold_old := 2
	shares_old := 5

	threshold_new := 10
	shares_new := 10

	size := 10
	secret := make([]byte, size)
	initialize(secret)

	splits, _ := shamir.SplitSecret(secret, shares_old, uint8(threshold_old))

	var redistShares [][]secretSharing.RedistShare

	for _, split := range splits {
		redist, err := redistribute.RedistributeShare(split, shares_new, threshold_new)
		if err != nil {
			println(err)
		}

		redistShares = append(redistShares, redist)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := redistribute.RedistSharesToShareholders(redistShares)
		if err != nil {
			println(err)
		}
	}
}

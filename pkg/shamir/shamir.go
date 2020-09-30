package shamir

import (
	"crypto/rand"
	"math"
	"math/big"
)

type share struct {
	shareIndex int
	slices     []int
}

const prime = 257

func SplitSecret(secret []byte, shares int, threshold int) []share {
	//TODO: Assertions...

	var sharedSecrets []share

	var indices []int

	for i := 1; i <= shares; i++ {
		indices = append(indices, i)
	}
	for _, byte := range secret {
		splitByte(byte, shares, threshold, indices)
	}
	return shared_secrets
}

func splitByte(secretByte byte, shares int, threshold int, indices []int) error {
	type byteShare struct {
		index       int
		coefficient int64
	}
	var byteshares []byteShare

	for _, index := range indices {
		coefficient, err := rand.Int(rand.Reader, big.NewInt(prime))
		if err != nil {
			return err
		}
		byteshares = append(byteshares, byteShare{index, coefficient.Int64()})
	}
	// Create Polynomial, the constant term is the secret, the maximum degree is threshold - 1
	var polynomial = func(x int) int {
		sum := uint64(secretByte)
		for i := 1; i < threshold; i++ {
			sum += uint64(byteshares[i-1].coefficient) * uint64(math.Pow(float64(x), float64(i)))
		}

		return sum % prime

	}
}

package shamir

import (
	"crypto/rand"
	"math"
	"math/big"
)

const prime = 257

type share struct {
	shareIndex int
	slices     []int
}
type singleByteShare struct {
	shareIndex int
	share      int
}

func SplitSecret(secret []byte, shares int, threshold int) ([]share, error) {
	//TODO: Assertions...

	var sharedSecrets []share
	var indices []int

	// initialize indices
	for i := 1; i <= shares; i++ {
		indices = append(indices, i)
	}

	// Split every byte
	secretSharesMap := make(map[int][]int)
	for _, byte := range secret {
		byteShares, err := splitByte(byte, shares, threshold, indices)
		if err != nil {
			return nil, err
		}

		for _, byteShare := range byteShares {
			// append to map
			bytes := secretSharesMap[byteShare.shareIndex]
			bytes = append(bytes, byteShare.share)
			secretSharesMap[byteShare.shareIndex] = bytes
		}
	}

	for index, bytes := range secretSharesMap {
		sharedSecrets = append(sharedSecrets, share{
			shareIndex: index,
			slices:     bytes,
		})
	}
	return sharedSecrets, nil
}

func splitByte(secretByte byte, shares int, threshold int, indices []int) ([]singleByteShare, error) {
	// Create Polynomial, the constant term is the secret, the maximum degree is threshold - 1
	polynom, err := buildRandomPolynomial(int(secretByte), threshold-1, prime)
	if err != nil {
		return nil, err
	}
	var singleByteShares []singleByteShare
	for index := range indices {
		singleByteShares = append(
			singleByteShares,
			singleByteShare{
				shareIndex: index,
				share:      polynom(index),
			})
	}
	return singleByteShares, nil
}

// Builds a polynomial function with a maximum degree, random coefficients, a
// given constant term and a given modulo.
func buildRandomPolynomial(constant, maxDegree, modulo int) (func(x int) int, error) {
	// first coefficient is the secret
	coefficients := []int{constant}
	for i := 0; i < maxDegree-1; i++ {
		coefficient, err := rand.Int(rand.Reader, big.NewInt(prime))
		if err != nil {
			return nil, err
		}

		coefficients = append(coefficients, int(coefficient.Int64()))
	}

	return buildPolynomial(coefficients, prime), nil
}

// Builds a polynomial function with given coefficients, and degree len(coefficients) + 1
func buildPolynomial(coefficents []int, modulo int) func(x int) int {
	f := func(x int) int {
		sum := 0
		for i, coefficient := range coefficents {
			sum += coefficient * int(math.Pow(float64(x), float64(i)))
		}
		return sum % modulo

	}
	return f
}

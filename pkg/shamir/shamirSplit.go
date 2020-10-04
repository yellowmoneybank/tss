package shamir

import (
	"crypto/rand"
	"github.com/google/uuid"
	"math"
	"math/big"
)

const prime = 257

type share struct {
	id         uuid.UUID
	threshold  uint8
	shareIndex uint16
	slices     []uint16
}
type singleByteShare struct {
	shareIndex uint16
	share      uint16
}

func SplitSecret(secret []byte, shares int, threshold uint8) ([]share, error) {
	//TODO: Assertions...

	var sharedSecrets []share
	var indices []uint16

	// initialize indices
	for i := 1; i <= shares; i++ {
		indices = append(indices, uint16(i))
	}

	// Split every byte
	// index => bytes
	secretSharesMap := make(map[uint16][]uint16)
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

	// all Shares have the same UUID
	uuid := uuid.New()
	for index, bytes := range secretSharesMap {
		sharedSecrets = append(sharedSecrets, share{
			id:         uuid,
			threshold:  threshold,
			shareIndex: index,
			slices:     bytes,
		})
	}
	return sharedSecrets, nil
}

func splitByte(secretByte byte, shares int, threshold uint8, indices []uint16) ([]singleByteShare, error) {
	// Create Polynomial, the constant term is the secret, the maximum degree is threshold - 1
	polynom, err := buildRandomPolynomial(int(secretByte), int(threshold-1), prime)
	if err != nil {
		return nil, err
	}
	var singleByteShares []singleByteShare
	for _, index := range indices {
		singleByteShares = append(
			singleByteShares,
			singleByteShare{
				shareIndex: uint16(index),
				share:      uint16(polynom(int(index))),
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

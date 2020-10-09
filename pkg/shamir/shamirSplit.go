package shamir

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"

	"github.com/google/uuid"
)

const prime = 257

type Share struct {
	Id         uuid.UUID
	threshold  uint8
	ShareIndex uint16
	Slices     []uint16
}

type singleByteShare struct {
	shareIndex uint16
	share      uint16
}

func SplitSecret(secret []byte, shares int, threshold uint8) ([]Share, error) {
	// TODO: Assertions...

	var sharedSecrets []Share

	// Split every secretByte
	// index => bytes
	secretSharesMap := make(map[uint16][]uint16)

	for i, secretByte := range secret {
		byteShares, err := splitByte(secretByte, shares, threshold)
		if err != nil {
			return nil, err
		}

		for _, byteShare := range byteShares {
			// append to map
			bytes := secretSharesMap[byteShare.shareIndex]
			bytes = append(bytes, byteShare.share)
			secretSharesMap[byteShare.shareIndex] = bytes
		}

		fmt.Printf("splitted: %d \r", i)
	}

	// all Shares have the same UUID
	uuId := uuid.New()
	for index, bytes := range secretSharesMap {
		sharedSecrets = append(sharedSecrets, Share{
			Id:         uuId,
			threshold:  threshold,
			ShareIndex: index,
			Slices:     bytes,
		})
	}
	return sharedSecrets, nil
}

func splitByte(secretByte byte, shares int, threshold uint8) ([]singleByteShare, error) {
	// Create Polynomial, the constant term is the secret, the maximum degree is threshold - 1
	polynomial, err := buildRandomPolynomial(int(secretByte), int(threshold-1), prime)
	if err != nil {
		return nil, err
	}

	// initialize indices
	indices := createIndices(shares)

	singleByteShares := createByteShares(indices, polynomial)
	return singleByteShares, nil
}

func createIndices(indexCount int) []uint16 {
	var indices []uint16
	for i := 1; i <= indexCount; i++ {
		indices = append(indices, uint16(i))
	}
	return indices
}

func createByteShares(indices []uint16, polynomial func(x int) int) []singleByteShare {
	var singleByteShares []singleByteShare
	for _, index := range indices {
		singleByteShares = append(
			singleByteShares,
			singleByteShare{
				shareIndex: index,
				share:      uint16(polynomial(int(index))),
			})
	}
	return singleByteShares
}

// Builds a polynomial function with a maximum degree, random coefficients, a
// given constant term and a given modulo.
func buildRandomPolynomial(constant, maxDegree, modulo int) (func(x int) int, error) {
	// first coefficient is the secret
	coefficients := []int{constant}
	for i := 0; i < maxDegree; i++ {
		coefficient, err := rand.Int(rand.Reader, big.NewInt(prime))
		if err != nil {
			return nil, err
		}

		coefficients = append(coefficients, int(coefficient.Int64()))
	}
	polynomial, err := buildPolynomial(coefficients, modulo)
	if err != nil {
		return nil, err
	}
	return polynomial, nil
}

// Builds a polynomial function with given coefficients, and degree len(coefficients) + 1
func buildPolynomial(coefficients []int, modulo int) (func(x int) int, error) {
	f := func(x int) int {
		sum := 0
		for i, coefficient := range coefficients {
			sum += coefficient * int(math.Pow(float64(x), float64(i)))
		}
		return sum % modulo
	}
	return f, nil
}

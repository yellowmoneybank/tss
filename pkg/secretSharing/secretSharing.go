package secretSharing

import (
	"crypto/rand"
	"math"
	"math/big"

	"github.com/google/uuid"
)

type Share struct {
	ID          uuid.UUID
	Threshold   uint8
	ShareIndex  uint16
	Secrets     []ByteShare
	Prime, Q, G int
}

type ByteShare struct {
	Share uint16
	// at contains g^s, g^a_1 ... for Feldman's VSS
	CheckValues []uint16
}

func RandomCoefficients(count, limit int) ([]int, error) {
	coefficients := make([]int, count)
	for i := 0; i < count; i++ {
		coefficient, err := rand.Int(rand.Reader, big.NewInt(int64(limit)))
		if err != nil {
			return nil, err
		}

		coefficients[i] = int(coefficient.Int64())
	}

	return coefficients, nil
}

// Builds a polynomial function with given coefficients, and degree len(coefficients) + 1.
func BuildPolynomial(coefficients []int, modulo int) (func(x int) int, error) {
	f := func(x int) int {
		sum := 0
		for i, coefficient := range coefficients {
			sum += coefficient * int(math.Pow(float64(x), float64(i)))
		}

		return sum % modulo
	}

	return f, nil
}

func BuildRandomPolynomial(constant, maxDegree, modulo int) (func(x int) int, error) {
	coefficients := []int{constant}

	randomCoefficients, err := RandomCoefficients(maxDegree, modulo)
	if err != nil {
		return nil, err
	}

	coefficients = append(coefficients, randomCoefficients...)

	polynomial, err := BuildPolynomial(coefficients, modulo)
	if err != nil {
		return nil, err
	}

	return polynomial, nil
}

func CreateIndices(indexCount int) []int {
	var indices []int
	for i := 1; i <= indexCount; i++ {
		indices = append(indices, i)
	}

	return indices
}

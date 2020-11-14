package shamir

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"moritzm-mueller.de/tss/pkg/feldman"
	"moritzm-mueller.de/tss/pkg/secretSharing"

	"github.com/google/uuid"
)

// p * r + 1 = q
// g is generator over Z_prime
// needed for Feldman's VSS.
const (
	p = 257
	//	r     = 6
	q = 1543
	g = 5
)


func SplitSecret(secret []byte, shares int, threshold uint8) ([]secretSharing.Share, error) {
	// TODO: Assertions...

	var sharedSecrets []secretSharing.Share

	// Split every secretByte, every index has multiple byteShares
	// index => byteShares
	secretSharesMap := make(map[uint16][]secretSharing.ByteShare)

	for i, secretByte := range secret {
		byteShares, err := splitByte(secretByte, shares, threshold)
		if err != nil {
			return nil, err
		}

		for index, byteShare := range byteShares {
			// append to map
			shares := secretSharesMap[uint16(index)]
			shares = append(shares, byteShare)
			secretSharesMap[uint16(index)] = shares
		}

		fmt.Printf("splitted: %d \r", i)
	}

	// all Shares have the same UUID
	uuID := uuid.New()

	for index, byteShares := range secretSharesMap {
		sharedSecrets = append(sharedSecrets, secretSharing.Share{
			ID:         uuID,
			Threshold:  threshold,
			ShareIndex: index,
			Secrets:    byteShares,
			Prime:      p,
			Q:          q,
			G:          g,
		})
	}

	return sharedSecrets, nil
}

// splits secretByte into shares with threshold. Return a map that maps the indices to the shares
func splitByte(secretByte byte, shares int, threshold uint8) (map[int]secretSharing.ByteShare, error) {
	// Create Polynomial, the constant term is the secret, the maximum degree is threshold - 1
	coefficients := []int{int(secretByte)}
	randomCoefficients, err := randomCoefficients(int(threshold - 1))
	if err != nil {
		return nil, err
	}

	coefficients = append(coefficients, randomCoefficients...)

	polynomial, err := buildPolynomial(coefficients, p)
	if err != nil {
		return nil, err
	}

	// initialize indices
	indices := createIndices(shares)

	singleByteShares := createByteShares(indices, polynomial)

	// initialize Checkvalues

	checkValues := feldman.CalculateCheckValues(g, q, coefficients)
	// type conversion
	checkvaluesUint16 := make([]uint16, len(checkValues))
	for _, value := range checkValues {
		checkvaluesUint16 = append(checkvaluesUint16, uint16(value))
	}

	for _, byteShare := range singleByteShares {

		copy(byteShare.CheckValues, checkvaluesUint16)
	}
	return singleByteShares, nil
}

func createIndices(indexCount int) []int {
	var indices []int
	for i := 1; i <= indexCount; i++ {
		indices = append(indices, i)
	}

	return indices
}

func createByteShares(indices []int, polynomial func(x int) int) map[int]secretSharing.ByteShare {
	singleByteShares := make(map[int]secretSharing.ByteShare)
	for _, index := range indices {
		singleByteShares[index] = secretSharing.ByteShare{
			Share:       uint16(polynomial(index)),
			CheckValues: nil,
		}
	}

	return singleByteShares
}

func randomCoefficients(number int) ([]int, error) {
	coefficients := make([]int, number)
	for i := 0; i < number; i++ {
		coefficient, err := rand.Int(rand.Reader, big.NewInt(p))
		if err != nil {
			return nil, err
		}

		coefficients[i] =  int(coefficient.Int64())
	}

	return coefficients, nil
}

// Builds a polynomial function with given coefficients, and degree len(coefficients) + 1.
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

package shamir

import (
	"moritzm-mueller.de/tss/pkg/feldman"

	"moritzm-mueller.de/tss/pkg/secretSharing"

	"github.com/google/uuid"
)

// p * r + 1 = q
// g is generator over Z_prime
// needed for Feldman's VSS.
const (
	p = 281
	//	r     = 2
	q = 563
	g = 11
)

func SplitSecret(secret []byte, shares int, threshold uint8) ([]secretSharing.Share, error) {
	var sharedSecrets []secretSharing.Share

	// Split every secretByte, every index has multiple byteShares
	// index => byteShares
	secretSharesMap := make(map[uint16][]secretSharing.ByteShare)

	for _, secretByte := range secret {
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

		// fmt.Printf("splitted: %d \r", i)
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

// splits secretByte into shares with threshold. Return a map that maps the indices to the shares.
func splitByte(secretByte byte, shares int, threshold uint8) (map[int]secretSharing.ByteShare, error) {
	// Create Polynomial, the constant term is the secret, the maximum degree is threshold - 1
	coefficients := []int{int(secretByte)}

	randomCoefficients, err := secretSharing.RandomCoefficients(int(threshold-1), p)
	if err != nil {
		return nil, err
	}

	coefficients = append(coefficients, randomCoefficients...)
	// coefficients := []int{int(secretByte), 222,32}
	polynomial, err := secretSharing.BuildPolynomial(coefficients, p)
	if err != nil {
		return nil, err
	}

	// initialize indices
	indices := secretSharing.CreateIndices(shares)

	singleByteShares := createByteShares(indices, polynomial)

	// initialize checkValues for Feldman
	checkValues := feldman.CalculateCheckValues(g, q, coefficients)
	// type conversion
	var checkValuesUint16 []uint16
	for _, value := range checkValues {
		checkValuesUint16 = append(checkValuesUint16, uint16(value))
	}

	for i, byteShare := range singleByteShares {
		byteShare.CheckValues = make([]uint16, len(checkValuesUint16))
		copy(byteShare.CheckValues, checkValuesUint16)
		// Checkvalue for  VSR
		byteShare.GS = checkValuesUint16[0]
		singleByteShares[i] = byteShare
	}

	return singleByteShares, nil
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

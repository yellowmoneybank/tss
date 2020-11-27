package redistribute

import (
	"moritzm-mueller.de/tss/pkg/feldman"
	"moritzm-mueller.de/tss/pkg/secretSharing"
)

func RedistributeShare(share secretSharing.Share, numberOfNewShares, newThreshold int) ([]secretSharing.RedistShare, error) {
	var redistShares []secretSharing.RedistShare

	splittedShares := make(map[int][]secretSharing.ByteShare)

	for _, secret := range share.Secrets {
		split, err := redistributeByteShare(secret, share.Prime, share.Q, share.G, numberOfNewShares, newThreshold)
		if err != nil {
			return nil, err
		}

		for index, secret := range split {
			splittedShares[index] = append(splittedShares[index], secret)
		}
	}

	for index, secrets := range splittedShares {
		redistShares = append(redistShares, secretSharing.RedistShare{
			OldIndex:     share.ShareIndex,
			NewIndex:     uint16(index),
			NewThreshold: uint8(newThreshold),
			Share: secretSharing.Share{
				ID:         share.ID,
				Threshold:  uint8(newThreshold),
				ShareIndex: uint16(index),
				Secrets:    secrets,
				Prime:      share.Prime,
				Q:          share.Q,
				G:          share.G,
			},
		})
	}

	return redistShares, nil
}

func redistributeByteShare(byteshare secretSharing.ByteShare, prime, q, g, shares, threshold int) (map[int]secretSharing.ByteShare, error) {
	coefficients := []int{int(byteshare.Share)}

	randomCoefficients, err := secretSharing.RandomCoefficients(threshold-1, prime)
	if err != nil {
		return nil, err
	}

	coefficients = append(coefficients, randomCoefficients...)

	polynomial, err := secretSharing.BuildPolynomial(coefficients, prime)
	if err != nil {
		return nil, err
	}

	indices := secretSharing.CreateIndices(shares)

	checkValues := feldman.CalculateCheckValues(g, q, coefficients)

	var checkValuesUint16 []uint16
	for _, value := range checkValues {
		checkValuesUint16 = append(checkValuesUint16, uint16(value))
	}

	// x => y
	newShares := make(map[int]secretSharing.ByteShare)
	for _, index := range indices {
		newShares[index] = secretSharing.ByteShare{
			Share:       uint16(polynomial(index)),
			GS:          byteshare.GS,
			CheckValues: checkValuesUint16,
		}
	}

	return newShares, nil
}

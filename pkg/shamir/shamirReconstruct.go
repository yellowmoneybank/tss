package shamir

import (
	"errors"
	"fmt"

	"moritzm-mueller.de/tss/pkg/secretSharing"
)

func Reconstruct(shares []secretSharing.Share) ([]byte, error) {
	// TODO Assertions...
	var secret []byte

	for i := 0; i < len(shares[0].Secrets); i++ {
		byteShares := make(map[int]secretSharing.ByteShare)
		for _, share := range shares {
			byteShares[int(share.ShareIndex)] = share.Secrets[i]
		}

		reconstructedByte, err := reconstructByte(byteShares, shares[0].Threshold, shares[0].Prime)
		if err != nil {
			return nil, err
		}

		secret = append(secret, reconstructedByte)

		// fmt.Printf("reconstructed: %d \r", i)
	}

	return secret, nil
}

func reconstructByte(byteShares map[int]secretSharing.ByteShare, threshold uint8, prime int) (byte, error) {
	if !secretSharing.IsUniqueSolution(byteShares, threshold) {
		return 0, errors.New("can't find unique solution")
	}

	var points []secretSharing.Point

	for x, byteShare := range byteShares {
		points = append(points, secretSharing.Point{X: x, Y: int(byteShare.Share)})
	}

	p := secretSharing.ReconstructPolynom(points, prime)

	secret := p(0)

	if secret != float64(int(secret)) {
		fmt.Println("We have a Problem")
	}

	return byte(secret), nil
}

package redistribute

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"moritzm-mueller.de/tss/pkg/secretSharing"
	"moritzm-mueller.de/tss/pkg/shamir"
)

func TestReconstruct(t *testing.T) {
	{
		secret := []byte{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}

		splits, _ := shamir.SplitSecret(secret, 3, 3)

		var redistShares [][]RedistShare

		for _, split := range splits {
			redist, err := RedistributeShare(split, 4, 4)
			if err != nil {
				println(err)
			}

			redistShares = append(redistShares, redist)
		}

		newShareholders, _ := redistSharesToShareholders(redistShares)

		reconstruct, _ := shamir.Reconstruct(newShareholders)
		assert.Equal(t, secret, reconstruct)
	}
}

func redistSharesToShareholders(redistShares [][]RedistShare) ([]secretSharing.Share, error) {
	orderdredistShares := make(map[uint16][]RedistShare)

	for _, redistShare := range redistShares {
		for _, indexShare := range redistShare {
			orderdredistShares[indexShare.newIndex] = append(orderdredistShares[indexShare.newIndex], indexShare)
		}
	}

	var newShares []secretSharing.Share

	for _, shares := range orderdredistShares {
		share, err := Reconstruct(shares)
		if err != nil {
			return nil, err
		}

		newShares = append(newShares, share)
	}

	return newShares, nil
}

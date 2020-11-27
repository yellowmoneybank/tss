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

		splits, _ := shamir.SplitSecret(secret, 5, 2)

		var redistShares [][]secretSharing.RedistShare

		for _, split := range splits {
			redist, err := RedistributeShare(split, 10, 4)
			if err != nil {
				println(err)
			}

			redistShares = append(redistShares, redist)
		}

		newShareholders, err := RedistSharesToShareholders(redistShares[0:2])
		if err != nil {
			println(err)
			return
		}

		reconstruct, _ := shamir.Reconstruct(newShareholders)
		assert.Equal(t, secret, reconstruct)
	}
}

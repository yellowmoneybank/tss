package shamir

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitSecret(t *testing.T) {
	{
		secret := "very secret secret"
		var threshold uint8 = 3
		numberOfShares := 5
		shares, err := SplitSecret([]byte(secret), numberOfShares, threshold)

		assert.Nil(t, err)

		assert.Equal(t, len(shares), numberOfShares)

		uuid := shares[0].id
		for _, share := range shares {
			assert.Equal(t, uuid, share.id, "All shares should have the same id")
		}

		for _, share := range shares {
			assert.Equal(t, threshold, share.threshold, "Al Shares should have the same threshold")
		}

		len_share := len(shares[0].slices)
		for _, share := range shares {
			assert.Equal(t, len_share, len(share.slices), "All shares should have the same length")
		}

	}
}

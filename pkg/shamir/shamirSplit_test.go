package shamir

import (
	"moritzm-mueller.de/tss/pkg/secretSharing"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitSecret(t *testing.T) {
	{
		secret := "very secret secret"
		var threshold uint8 = 3
		numberOfShares := 5
		shares, err := SplitSecret([]byte(secret), numberOfShares, threshold)

		assert.Nil(t, err)

		assert.Equal(t, len(shares), numberOfShares)

		uuid := shares[0].ID
		for _, share := range shares {
			assert.Equal(t, uuid, share.ID, "All shares should have the same id")
		}

		for _, share := range shares {
			assert.Equal(t, threshold, share.Threshold, "Al Shares should have the same threshold")
		}

		lenShare := len(shares[0].Secrets)
		for _, share := range shares {
			assert.Equal(t, lenShare, len(share.Secrets), "All shares should have the same length")
		}
	}
}

func Test_buildPolynomial(t *testing.T) {
	poly, _ := secretSharing.BuildPolynomial([]int{10, 45, 102}, p)

	shouldBePolynomial := func(x int) int {
		return (10 + 45*x + 102*x*x) % p
	}

	for i := 1; i <= 3; i++ {
		assert.Equal(t, shouldBePolynomial(i), poly(i), "should be Equal")
	}
}

func Test_createByteShares(t *testing.T) {
	var indices []int
	for i := 1; i <= 5; i++ {
		indices = append(indices,i)
	}

	polynomial := func(x int) int {
		return (10 + 45*x + 102*x*x) % p
	}

	byteShares := createByteShares(indices, polynomial)

	for i, share := range byteShares {
		assert.Contains(t,indices, i)
		assert.Equal(t, polynomial(i), int(share.Share), "should be equal")
	}
}

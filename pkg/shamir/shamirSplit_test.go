package shamir

import (
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

		uuid := shares[0].id
		for _, share := range shares {
			assert.Equal(t, uuid, share.id, "All shares should have the same id")
		}

		for _, share := range shares {
			assert.Equal(t, threshold, share.threshold, "Al Shares should have the same threshold")
		}

		lenShare := len(shares[0].Slices)
		for _, share := range shares {
			assert.Equal(t, lenShare, len(share.Slices), "All shares should have the same length")
		}
	}
}

func Test_buildPolynomial(t *testing.T) {
	p, _ := buildPolynomial([]int{10, 45, 102}, prime)

	shouldBePolynomial := func(x int) int {
		return (10 + 45*x + 102*x*x) % prime
	}

	for i := 1; i <= 3; i++ {
		assert.Equal(t, shouldBePolynomial(i), p(i), "should be Equal")
	}
}

func Test_createByteShares(t *testing.T) {
	var indices []uint16
	for i := 1; i <= 5; i++ {
		indices = append(indices, uint16(i))
	}

	polynomial := func(x int) int {
		return (10 + 45*x + 102*x*x) % prime
	}

	byteShares := createByteShares(indices, polynomial)

	for i, share := range byteShares {
		assert.Equal(t, indices[i], share.shareIndex, "should be equal")
		assert.Equal(t, polynomial(int(indices[i])), int(share.share), "should be equal")
	}
}

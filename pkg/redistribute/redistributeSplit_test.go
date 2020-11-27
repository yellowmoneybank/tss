package redistribute

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"moritzm-mueller.de/tss/pkg/shamir"
)

func TestRedistributeShare(t *testing.T) {
	secret := "very secret secret"
	var threshold uint8 = 3
	numberOfShares := 5
	shares, _ := shamir.SplitSecret([]byte(secret), numberOfShares, threshold)

	numberOfNewShares := 6
	newThreshold := 5
	redistShares, err := RedistributeShare(shares[0], numberOfNewShares, newThreshold)

	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, numberOfNewShares, len(redistShares), "numberOfNewShares is wrong")
}

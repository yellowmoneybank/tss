package shamir

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestSplitSecret
func TestSplitSecret(t *testing.T) {

	shares := SplitSecret([]byte("secret"), 10, 3)

	for _, share := range shares {
		assert.GreaterOrEqual(t, share.shareIndex, 1)
	}

}

package cheaterDetection

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calcSum1(t *testing.T) {
	assert.NotEmpty(t, calcSum1(map[int]big.Int{
		1:  *big.NewInt(12),
		12: *big.NewInt(12),
		14: *big.NewInt(12),
	}, big.NewInt(13)))

	hashes := map[int]big.Int{
		3: *big.NewInt(1),
		1: *big.NewInt(7),
		2: *big.NewInt(3),
	}
	assert.Equal(t, big.NewInt(15011), calcSum1(hashes, big.NewInt(11)))
}

func Test_calcSum2(t *testing.T) {
	{
		assert.Equal(t, int64(1050), calcSum2(big.NewInt(3), big.NewInt(7), 3).Int64())
	}
	{
		assert.Equal(t, int64(4026), calcSum2(big.NewInt(3), big.NewInt(11), 3).Int64())
	}
}
